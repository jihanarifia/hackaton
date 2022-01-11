package server

import (
	"fmt"
	"net"
	"net/http"

	"hackaton/pkg/config"
	"hackaton/pkg/service"

	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"
)

const (
	ServiceName = "ABHack"
	versionPath = "/version"
)

// Server provide state and functionality to run REST API server
type Server struct {
	container *restful.Container
	listener  net.Listener
}

// New creates new Server instances based on configuration
func New(config *config.Config, service *service.Service) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", config.Port))
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return nil, err
	}

	server := &Server{
		listener:  listener,
		container: restful.NewContainer(),
	}

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders: []string{},
		AllowedHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Methods",
			"Content-Type",
			"Accept",
			"Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      server.container,
	}

	server.container.Filter(cors.Filter)

	rootWebservice := new(restful.WebService)
	server.AddVersionEndpoint(rootWebservice, service, config.BasePath)
	server.container.Add(rootWebservice)

	return server, nil
}

// Stop stops HTTP listener (REST API server)
func (server *Server) Stop() {
	err := server.listener.Close()
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}

// Serve starts HTTP listener for REST API server
func (server *Server) Serve() {
	err := http.Serve(server.listener, server.container)
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}

// AddVersionEndpoint add version endpoint that returns information about the service, name, build date and revision ID
func (server *Server) AddVersionEndpoint(webService *restful.WebService, service *service.Service, basePath string) {
	webService.Route(webService.GET(basePath + versionPath).To(service.VersionHandlerShort))
}
