package middleware

import (
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	"github.com/arymaulanamalik/bobobox_test/pkg/response"
)

// Recovery is a middleware that recovers from panics, logs the panic
// and returns a HTTP 500 (Internal Server Error) status.
func Recovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Warning("RECOVERY PANIC ERROR", err)
				response.Error(res, http.StatusInternalServerError, "There was an internal server error")
			}
		}()
		h.ServeHTTP(res, req)
	})
}
