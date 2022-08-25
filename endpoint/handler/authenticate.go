package handler

import (
	"context"
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/model"
	"github.com/arymaulanamalik/bobobox_test/pkg/response"
	"github.com/arymaulanamalik/bobobox_test/service"
)

type AuthenticateHandler struct {
	service service.AuthenticateService
}

func NewAuthenticateHandler(service service.AuthenticateService) *AuthenticateHandler {
	return &AuthenticateHandler{
		service: service,
	}
}

func (ah *AuthenticateHandler) Authenticate(res http.ResponseWriter, req *http.Request) {

	reasons, err := ah.service.Authenticate(context.Background(), model.AuthenticateRequest{})
	if err != nil {
		// response.Error(res, error.Code, error.Message)
		return
	}

	response.Success(res, http.StatusOK, reasons, nil)
}
