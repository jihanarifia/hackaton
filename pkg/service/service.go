package service

import (
	"github.com/emicklei/go-restful"
	"hackaton/pkg/config"
	"hackaton/version"
	"net/http"
)

type Service struct {
	ServiceName string
	Config      config.Config
}

func New(serviceName string, config config.Config) *Service {
	return &Service{
		ServiceName: serviceName,
		Config:      config,
	}
}

func (service *Service) VersionHandlerShort(request *restful.Request, response *restful.Response) {
	versionData := version.Short(service.ServiceName)
	err := response.WriteAsJson(&versionData)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}
