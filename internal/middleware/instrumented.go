package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/joepk90/graphql-auth/internal/logger"
	"github.com/joepk90/graphql-auth/internal/stats"
)

// HTTPInstrumentedInterceptor instruments the request
func HTTPInstrumentedInterceptor(next http.Handler, metrics stats.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log := logger.ForensicLoggerFromRequest(r)
		if r.Body != nil {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Warn("error reading from request body pre-proxy")
			}
			r.Body = io.NopCloser(bytes.NewBuffer(b))

			// Request stats
			if len(b) >= 0 {
				metrics.ObserveRequestSize(float64(len(b)))
			}
		}

		metrics.ObserveRequest()

		crw := CustomResponseWriter{ResponseWriter: w}
		next.ServeHTTP(&crw, r)

		// Response stats
		if crw.length > 0 {
			metrics.ObserveResponseSize(float64(crw.length))
		}
		metrics.ObserveResponseSize(float64(crw.length))
		metrics.ObserveResponseTime(time.Since(start))
	})
}
