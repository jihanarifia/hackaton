package service

import (
	"net/http"

	"hackaton/pkg/model"

	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
)

func (service *Service) GetUsers(req *restful.Request, resp *restful.Response) {
	user, err := service.DB.GetUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			responseErr(resp, 200, "user not found", nil)
			return
		}

		responseErr(resp, http.StatusInternalServerError, "internal server error", err)
	}

	responseSuccess(resp, 200, user)
}

func (service *Service) CreateUser(req *restful.Request, resp *restful.Response) {
	var userReq model.User
	if err := req.ReadEntity(&userReq); err != nil {
		responseErr(resp, http.StatusBadRequest, "Unable to parse request body", err)
		return
	}

	err := service.DB.CreateUser(userReq)
	if err != nil {
		responseErr(resp, http.StatusInternalServerError, "Unable to create new user", err)
		return
	}

	responseSuccess(resp, http.StatusCreated, userReq)
}
