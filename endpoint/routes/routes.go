package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	chitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"

	"github.com/arymaulanamalik/bobobox_test/endpoint/handler"
	"github.com/arymaulanamalik/bobobox_test/endpoint/middleware"
	"github.com/arymaulanamalik/bobobox_test/endpoint/registry"
	"github.com/arymaulanamalik/bobobox_test/repository"
	"github.com/arymaulanamalik/bobobox_test/service"
)

// GetRouter returns a Mux object that implements the Router interface.
func GetRouter(service service.AuthenticateService, repository repository.Repository) *chi.Mux {

	router := chi.NewRouter()

	setupMiddleware(router)

	healthCheckHandler := handler.NewHealthCheckHandler(repository)
	router.Get("/health", healthCheckHandler.HealthCheck)

	authenticateRegistry := registry.RegisterAuthenticate(service)
	// API version 1
	router.Route("/v1", func(router chi.Router) {
		// Add app authorization middleware
		router.Use(middleware.AppAuthorization)

		// register your routes here
		router.Mount("/auth", authenticateRegistry)
	})

	return router

}

// Setup global middlewares
func setupMiddleware(router *chi.Mux) {

	// Add a own middleware.
	router.Use(
		setCorsOptions(),
		middleware.Recovery,
		middleware.TraceIDHeader,
	)

	// Add a chi middleware.
	router.Use(
		chiMiddleware.Timeout(60*time.Second),
		chiMiddleware.StripSlashes,
		chitrace.Middleware(),
	)

}

func setCorsOptions() func(next http.Handler) http.Handler {

	// Basic CORS (cross-origin-resource-sharing)
	// for detail about CORS, see: https://github.com/rs/cors
	var corsOptions = cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Api-Key", "trace_id"},
		AllowCredentials: true,
		MaxAge:           300, // results of a preflight request can be cached in 5 minutes.
	})

	return corsOptions

}
