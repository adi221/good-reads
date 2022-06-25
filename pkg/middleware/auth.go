package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/helper"
)

// Authenticate authenticates the JWT token from req.cookie
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			// Supposed to return 401 but for now will leave it because there are routes that doesn't require authorization (login, signup)
			next.ServeHTTP(w, r)
			return
		}
		claims, err := helper.ExtractClaimsFromToken(authHeader[7:])
		if err != nil {
			jsonErrors(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// TODO: find user by claims["id"] and then save it in context + id
		ctx = context.WithValue(ctx, constant.ContextUserID, claims["id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
