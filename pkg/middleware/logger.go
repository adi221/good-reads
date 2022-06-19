package middleware

import (
	"net/http"
	"time"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/rs/zerolog/log"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			requestID, ok := r.Context().Value(constant.ContextRequestID).(string)
			if !ok {
				requestID = "unknown"
			}
			log.Debug().Str(
				"req-id", requestID,
			).Str(
				"remote-addr", r.RemoteAddr,
			).Str(
				"user-agent", r.UserAgent(),
			).Dur(
				"duration", time.Since(start),
			).Msgf("%s %s", r.Method, r.URL.Path)
		}()
		next.ServeHTTP(w, r)
	})
}
