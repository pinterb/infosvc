package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// EncodeClusterRequest encodes the request to the provided HTTP request, simply
// by JSON encoding to the request body. It's designed to be used in
// transport/http.Client.
func EncodeClusterRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// DecodeClusterRequest decodes the request from the provided HTTP request, simply
// by JSON decoding from the request body. It's designed to be used in
// transport/http.Server.
func DecodeClusterRequest(r *http.Request) (interface{}, error) {
	var request ClusterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeClusterResponse encodes the response to the provided HTTP response
// writer, simply by JSON encoding to the writer. It's designed to be used in
// transport/http.Server.
func EncodeClusterResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeClusterResponse decodes the response from the provided HTTP response,
// simply by JSON decoding from the response body. It's designed to be used in
// transport/http.Client.
func DecodeClusterResponse(resp *http.Response) (interface{}, error) {
	var response ClusterResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// EncodeHostRequest encodes the request to the provided HTTP request, simply
// by JSON encoding to the request body. It's designed to be used in
// transport/http.Client.
func EncodeHostRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// DecodeHostRequest decodes the request from the provided HTTP request, simply
// by JSON decoding from the request body. It's designed to be used in
// transport/http.Server.
func DecodeHostRequest(r *http.Request) (interface{}, error) {
	var request HostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeHostResponse encodes the response to the provided HTTP response
// writer, simply by JSON encoding to the writer. It's designed to be used in
// transport/http.Server.
func EncodeHostResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeHostResponse decodes the response from the provided HTTP response,
// simply by JSON decoding from the response body. It's designed to be used in
// transport/http.Client.
func DecodeHostResponse(resp *http.Response) (interface{}, error) {
	var response HostResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// EncodeHealthRequest encodes the request to the provided HTTP request, simply
// by JSON encoding to the request body. It's designed to be used in
// transport/http.Client.
func EncodeHealthRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// DecodeHealthRequest decodes the request from the provided HTTP request, simply
// by JSON decoding from the request body. It's designed to be used in
// transport/http.Server.
func DecodeHealthRequest(r *http.Request) (interface{}, error) {
	var request HealthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeHealthResponse encodes the response to the provided HTTP response
// writer, simply by JSON encoding to the writer. It's designed to be used in
// transport/http.Server.
func EncodeHealthResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeHealthResponse decodes the response from the provided HTTP response,
// simply by JSON decoding from the response body. It's designed to be used in
// transport/http.Client.
func DecodeHealthResponse(resp *http.Response) (interface{}, error) {
	var response HealthResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// EncodeHelloRequest encodes the request to the provided HTTP request, simply
// by JSON encoding to the request body. It's designed to be used in
// transport/http.Client.
func EncodeHelloRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// DecodeHelloRequest decodes the request from the provided HTTP request, simply
// by JSON decoding from the request body. It's designed to be used in
// transport/http.Server.
func DecodeHelloRequest(r *http.Request) (interface{}, error) {
	var request HelloRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeHelloResponse encodes the response to the provided HTTP response
// writer, simply by JSON encoding to the writer. It's designed to be used in
// transport/http.Server.
func EncodeHelloResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeHelloResponse decodes the response from the provided HTTP response,
// simply by JSON decoding from the response body. It's designed to be used in
// transport/http.Client.
func DecodeHelloResponse(resp *http.Response) (interface{}, error) {
	var response HelloResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
