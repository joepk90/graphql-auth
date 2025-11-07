package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// CORSHTTPMiddleware provides CORS support
func CORSHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Add("Access-Control-Allow-Headers", "content-type, authorization")

		if r.Method == http.MethodOptions {
			log.Debug("method == options")
			return
		}

		next.ServeHTTP(w, r)
	})
}
