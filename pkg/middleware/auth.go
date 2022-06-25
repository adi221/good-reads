package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/helper"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/adi221/good-reads/pkg/util"
)

// Authenticate authenticates the JWT token from req.cookie
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			// Graphql route contain routes that don't require authorization first (login, signup). So decision will be made inside the route.
			if r.URL.Path == "/graphql" {
				next.ServeHTTP(w, r)
				return
			}

			util.JsonErrors(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := helper.ExtractClaimsFromToken(authHeader[7:])
		if err != nil {
			util.JsonErrors(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		uid := uint(claims["id"].(float64))
		user, err := service.Lookup().GetUserById(ctx, uid)
		if err != nil {
			util.JsonErrors(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			util.JsonErrors(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, constant.ContextUser, *user)
		ctx = context.WithValue(ctx, constant.ContextUserID, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
