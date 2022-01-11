package service

import (
	"net/http"

	"hackaton/pkg/config"
	"hackaton/pkg/dao"
	"hackaton/pkg/model"
	"hackaton/version"

	"github.com/emicklei/go-restful"
)

type Service struct {
	ServiceName string
	DB          dao.DB
	Config      config.Config
}

func New(serviceName string, db dao.DB, config config.Config) *Service {
	return &Service{
		ServiceName: serviceName,
		DB:          db,
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

func (service *Service) HealthCheckHandler(req *restful.Request, resp *restful.Response) {
	dependencies := make([]model.Dependency, 0)

	postgresDependency := model.Dependency{
		Name:    "postgres",
		Healthy: true,
	}
	if err := service.DB.Health(); err != nil {
		postgresDependency.Healthy = false
	}
	dependencies = append(dependencies, postgresDependency)

	healthcheckResponse := model.HealthCheckResponse{
		Name:         service.ServiceName,
		Dependencies: dependencies,
		Healthy:      true,
	}

	for _, val := range dependencies {
		if !val.Healthy {
			healthcheckResponse.Healthy = false
			responseSuccess(resp, http.StatusServiceUnavailable, healthcheckResponse)
			return
		}
	}

	healthcheckResponse.Healthy = true
	responseSuccess(resp, http.StatusOK, healthcheckResponse)
}
