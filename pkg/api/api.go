package api

import (
	"net/http"

	"hackaton/pkg/model"
	"hackaton/pkg/service"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

func AddUserRoute(service *service.Service, basePath string) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path(basePath + "/user").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Doc("User")

	tags := []string{"User"}

	webService.Route(webService.GET("/").To(service.GetUsers).
		Notes("Get All Users").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), model.SuccessResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), model.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), model.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get All Users"))

	webService.Route(webService.POST("/").To(service.CreateUser).
		Reads(model.User{}).
		Notes("Create User").
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), model.SuccessResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), model.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), model.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Create new User"))

	return webService
}
