package api

import (
	"net/http"

	"github.com/adi221/good-reads/pkg/enum"
	"github.com/adi221/good-reads/pkg/middleware"
)

// Route is the structure of an HTTP route definition
type Route struct {
	Path        string
	Handler     http.Handler
	Middlewares []middleware.Middleware
}

func newRoute(path string, handler http.Handler, middlewares ...middleware.Middleware) Route {
	return Route{
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}
}

// Routes is a list of Route
type Routes []Route

func routes() Routes {
	return Routes{
		newRoute(
			"/",
			index(),
			middleware.Methods(enum.HttpMethods.GET),
		),
		newRoute(
			"/graphql",
			graphqlHandler(),
			middleware.RequestContext,
			middleware.Authenticate,
			middleware.Methods(enum.HttpMethods.GET, enum.HttpMethods.POST),
		),
	}
}
