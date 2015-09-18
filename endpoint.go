package main

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pinterb/infosvc/server"
)

func makeHelloEndpoint(svc server.InfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(server.HelloRequest)
		saying, err := svc.Hello(req.Name)
		var errText string
		if err != nil {
		    errText = err.Error() 
        }
		return server.HelloResponse{Saying: saying, Err: errText}, nil
	}
}

func makeHostEndpoint(svc server.InfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		host, err := svc.Host()
		return host, err
	}
}
