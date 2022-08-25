package registry

import (
	"github.com/arymaulanamalik/bobobox_test/endpoint/handler"
	"github.com/arymaulanamalik/bobobox_test/service"
	"github.com/go-chi/chi"
)

func RegisterAuthenticate(service service.AuthenticateService) *chi.Mux {
	authenticateHandler := handler.NewAuthenticateHandler(service)

	routes := chi.NewRouter()

	routes.Group(func(router chi.Router) {
		routes.Get("/login", authenticateHandler.Authenticate)
	})

	return routes

}
