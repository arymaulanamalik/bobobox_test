package service

import (
	"context"

	"github.com/arymaulanamalik/bobobox_test/model"
)

type AuthenticateService interface {
	Authenticate(ctx context.Context, req model.AuthenticateRequest) (model.AuthenticateResponse, error)
}
