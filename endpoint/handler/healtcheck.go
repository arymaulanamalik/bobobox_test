package handler

import (
	"log"
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/pkg/response"
	"github.com/arymaulanamalik/bobobox_test/repository"
)

type HealthCheckHandler struct {
	repository repository.Repository
}

func NewHealthCheckHandler(repository repository.Repository) *HealthCheckHandler {
	return &HealthCheckHandler{
		repository: repository,
	}
}

func (hc *HealthCheckHandler) HealthCheck(res http.ResponseWriter, req *http.Request) {
	log.Println("handler healthcheck")
	connect := hc.repository.ReadWriter.VerifyConnection()
	if !connect {
		// response.Error(res, error.Code, error.Message)
		return
	}

	response.Success(res, http.StatusOK, "", nil)
}
