package main  // import "github.com/pinterb/infosvc"

import (
	"flag"
	"fmt"
	stdlog "log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
	"github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	
	"github.com/pinterb/infosvc/server"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		debugAddr                    = fs.String("debug.addr", ":8000", "Address for HTTP debug/instrumentation server")
		httpAddr                     = fs.String("http.addr", ":8001", "Address for HTTP (JSON) server")
	)

	flag.Usage = fs.Usage // only show our flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// package log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
		stdlog.SetFlags(0)                             // flags are handled by Go kit's logger
		stdlog.SetOutput(log.NewStdlibAdapter(logger)) // redirect anything using stdlib log to us
	}

	// package metrics
	var requestDuration metrics.TimeHistogram
	{
		requestDuration = metrics.NewTimeHistogram(time.Nanosecond, metrics.NewMultiHistogram(
			expvar.NewHistogram("request_duration_ns", 0, 5e9, 1, 50, 95, 99),
			prometheus.NewSummary(stdprometheus.SummaryOpts{
				Namespace: "myorg",
				Subsystem: "addsvc",
				Name:      "duration_ns",
				Help:      "Request duration in nanoseconds.",
			}, []string{"method"}),
		))
	}
    
	// Business domain
	var svc server.InfoService
	{
		svc = pureInfoService{}
		svc = loggingMiddleware{svc, logger}
		svc = instrumentingMiddleware{svc, requestDuration}
	}

	// Mechanical stuff
	rand.Seed(time.Now().UnixNano())
	root := context.Background()
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	// Debug/instrumentation
	go func() {
		transportLogger := log.NewContext(logger).With("transport", "debug")
		_ = transportLogger.Log("addr", *debugAddr)
		errc <- http.ListenAndServe(*debugAddr, nil) // DefaultServeMux
	}()

	// Transport: HTTP/JSON
	go func() {
		var (
			transportLogger = log.NewContext(logger).With("transport", "HTTP/JSON")
			mux             = http.NewServeMux()
			hello, host      endpoint.Endpoint
		)

		hello = makeHelloEndpoint(svc)
		mux.Handle("/hello", httptransport.NewServer(
			root,
			hello,
			server.DecodeHelloRequest,
			server.EncodeHelloResponse,
			httptransport.ServerErrorLogger(transportLogger),
		))

		host = makeHostEndpoint(svc)
		mux.Handle("/host", httptransport.NewServer(
			root,
			host,
			server.DecodeHostRequest,
			server.EncodeHostResponse,
			httptransport.ServerErrorLogger(transportLogger),
		))

		_ = transportLogger.Log("addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, mux)
	}()

	_ = logger.Log("fatal", <-errc) 
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
