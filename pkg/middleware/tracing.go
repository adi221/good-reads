package middleware

import (
	"context"
	"net/http"

	"github.com/adi221/good-reads/pkg/constant"
)

// Tracing is a middleware to trace HTTP request
func Tracing(nextRequestID func() string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), constant.ContextRequestID, requestID)
			w.Header().Set("X-Request-ID", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
