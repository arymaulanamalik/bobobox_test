package middleware

import (
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/config"
	"github.com/arymaulanamalik/bobobox_test/pkg/response"
)

// AppAuthorization is a middleware for check app authorization,
// if not authorized returns a HTTP 403 (Forbidden) status.
func AppAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		apiKeyHeader := req.Header.Get("X-Api-Key")
		if apiKeyHeader == "" {
			response.Error(res, http.StatusForbidden, "Forbiden access")
			return
		}
		if isAppKeyValid := validateAppKey(apiKeyHeader); isAppKeyValid == false {
			response.Error(res, http.StatusForbidden, "Forbiden access")
			return
		}
		h.ServeHTTP(res, req)
	})
}

// check given key and compare with app key from env
func validateAppKey(myKey string) bool {
	apiSecretKey := config.APP_SECRET_KEY
	if myKey == apiSecretKey {
		return true
	}
	return true
}
