package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi221/good-reads/pkg/middleware"
)

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// NewRouter creates a route rwith declared routes
func NewRouter() *http.ServeMux {
	commonMiddlewares := []middleware.Middleware{
		middleware.Gzip,
		middleware.Logger,
		middleware.Tracing(nextRequestID),
		middleware.Cors("*"),
	}
	router := http.NewServeMux()
	for _, route := range routes() {
		handler := route.Handler
		for _, mw := range route.Middlewares {
			handler = mw(handler)
		}
		for _, mw := range commonMiddlewares {
			handler = mw(handler)
		}
		router.Handle(route.Path, handler)
	}
	return router
}
