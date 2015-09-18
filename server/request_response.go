package server

// HelloRequest is the business domain type for a Hello method request.
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse is the business domain type for a Hello method response.
type HelloResponse struct {
	Saying string `json:"saying"`
	Err string `json:"err,omitempty"`
}

// HostRequest is the business domain type for a Host method request.
type HostRequest struct {
	Name string `json:"name"`
}

// HostResponse is the business domain type for a Host method response.
type HostResponse struct {
	Name string `json:"name"` // https://golang.org/src/os/os_test.go
	Addrs []string `json:"addresses"` // http://play.golang.org/p/Apqck0ovcr 
	OS string `json:"os"` // Env Variable Lookup -- https://golang.org/src/os/env.go?s=2368:2398#L69
	Err string `json:"err,omitempty"`
}

// HealthRequest is the business domain type for a Health method request.
type HealthRequest struct {
	Details bool `json:"details"`
}

// HealthResponse is the business domain type for a Health method response.
type HealthResponse struct {
	Status string `json:"status"`
}

// ClusterRequest is the business domain type for a Cluster method request.
type ClusterRequest struct {
	Name string `json:"name"`
}

// ClusterResponse is the business domain type for a Cluster method response.
type ClusterResponse struct {
	Name string `json:"name"` // https://golang.org/src/os/os_test.go
	Nodes []string `json:"nodes"` // could be / should be a 'node' struct
	Services []string `json:"services"` // could be / should be a 'service' struct
	Err string `json:"err,omitempty"`
}
