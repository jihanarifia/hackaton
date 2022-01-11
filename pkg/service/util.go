package service

import (
	"encoding/json"

	"hackaton/pkg/model"

	"github.com/emicklei/go-restful"
)

func writeResponse(resp *restful.Response, statusCode int, entity interface{}) error {
	resp.Header().Set("Content-Type", restful.MIME_JSON)
	resp.WriteHeader(statusCode)
	enc := json.NewEncoder(resp)
	enc.SetEscapeHTML(false)
	return enc.Encode(entity)
}

func responseSuccess(resp *restful.Response, statusCode int, entity interface{}) {
	data := model.SuccessResponse{
		Data: entity,
	}

	writeResponse(resp, statusCode, data)
}

func responseErr(resp *restful.Response, statusCode int, message string, err interface{}) {
	data := model.ErrorResponse{
		Message: message,
		Error:   err,
	}

	writeResponse(resp, statusCode, data)
}
