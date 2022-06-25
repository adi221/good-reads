package middleware

import (
	"context"
	"net/http"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/helper"
)

// RequestContext adds additional info to the request context
func RequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, constant.ContextCookieSetter, func(cookieValue string) {
			helper.SetCookieHandler(w, cookieValue)
		})
		ctx = context.WithValue(ctx, constant.ContextAuthorizationTokenSetter, func(stringToken string) {
			helper.AddTokenToResponseHeader(w, stringToken)
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
