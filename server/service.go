package server

// InfoService is the abstract representation of this service.
type InfoService interface {
	Host() (HostResponse, error)
//	Health() (string)
//	Cluster(string) (string, error)
	Hello(string) (string, error)
}
