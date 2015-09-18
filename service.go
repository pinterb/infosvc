package main

import (
    "fmt"
	"time"
	"net"

    "github.com/kelseyhightower/envconfig"
	"github.com/pinterb/infosvc/server"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

type pureInfoService struct{}

func (pureInfoService) Hello(name string) (saying string, err error) {
    saying = fmt.Sprintf("Hello %s! Hope all is well...", name)
    return saying, nil
}

// HostSec represents the list of environment variables to expect for the Host endpoint. 
type HostSpec struct {
    Host string
    OS   string
    Arch string
}

func (pureInfoService) Host() (server.HostResponse, error) {
	var errResp string
	var hostSpec HostSpec
	
	err := envconfig.Process("infosvc", &hostSpec)
	if err != nil {
	    errResp = err.Error() 
    }

    var ips []string
    
   	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

    var myOS = fmt.Sprintf("%s %s", hostSpec.OS, hostSpec.Arch)
	return server.HostResponse{Name: hostSpec.Host, OS: myOS, Addrs: ips, Err: errResp}, nil
}

type loggingMiddleware struct {
	server.InfoService
	log.Logger
}

func (m loggingMiddleware) Hello(name string) (saying string, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "hello",
			"name", name,
			"saying", saying,
			"took", time.Since(begin),
		)
	}(time.Now())

	saying, err = m.InfoService.Hello(name)
	return
}

func (m loggingMiddleware) Host() (host server.HostResponse, err error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "host",
			"name", host.Name,
			"os", host.OS,
			"took", time.Since(begin),
		)
	}(time.Now())
	
	host, err = m.InfoService.Host()
	return
}

type instrumentingMiddleware struct {
	server.InfoService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Hello(name string) (saying string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "hello"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())

	saying, err = m.InfoService.Hello(name)
	return
}

func (m instrumentingMiddleware) Host() (host server.HostResponse, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "host"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	
	host, err = m.InfoService.Host()
	return
}
