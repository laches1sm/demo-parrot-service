package httpserver

import (
	"../../adapters"
	"log"
	"net/http"
)

const (
	ServerPort     = ":7000"
	ParrotEndpoint = "/parrots"
)

// ParrotServer is an interface to an HTTP server which handles requests for parrots
type ParrotServer struct {
	Mux               *http.ServeMux
	Logger            log.Logger
	ParrotHTTPAdapter *adapters.ParrotHTTPAdapter
}

// SetupRoutes configures the routes of the API
func (srv *ParrotServer) SetupRoutes() {
	srv.Mux.Handle(ParrotEndpoint, http.HandlerFunc(srv.ParrotHTTPAdapter.GibeParrot))
	srv.Mux.Handle(ParrotEndpoint, http.HandlerFunc(srv.ParrotHTTPAdapter.AddParrot))
}

// Start sets up the HTTP webserver to listen and handle traffic. It
// takes the port number to listen on as a parameter in the form ":PORT_NUMBER"
func (srv *ParrotServer) Start(port string) error {
	return http.ListenAndServe(port, srv.Mux)
}

// NewParrotServer returns an instance of a configured ParrotServer
func NewParrotServer(logger log.Logger, adapter *adapters.ParrotHTTPAdapter) *ParrotServer {
	httpServer := &ParrotServer{
		Mux:               http.NewServeMux(),
		Logger:            logger,
		ParrotHTTPAdapter: adapter,
	}
	return httpServer
}
