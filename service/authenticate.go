package service

import (
	"context"

	"github.com/arymaulanamalik/bobobox_test/model"
	"github.com/arymaulanamalik/bobobox_test/repository"
)

type authenticate struct {
	Repository repository.Repository
}

func NewAuthenticateService(repository repository.Repository) AuthenticateService {
	return &authenticate{
		Repository: repository,
	}
}

func (a *authenticate) Authenticate(context.Context, model.AuthenticateRequest) (model.AuthenticateResponse, error) {
	return model.AuthenticateResponse{}, nil
}
