package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

// TraceIDHeader is a middleware that injects a request Header trace_id
// and set a response header
func TraceIDHeader(next http.Handler) http.Handler {

	traceIDHeader := "trace_id"

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		traceID := req.Header.Get(traceIDHeader)
		if traceID == "" {
			randomTraceID := uuid.New()
			traceID = randomTraceID.String()

			req.Header.Add(traceIDHeader, traceID)
		}

		res.Header().Set(traceIDHeader, traceID)

		next.ServeHTTP(res, req)
	})
}
