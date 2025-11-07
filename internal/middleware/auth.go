package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	iam "github.com/joepk90/graphql-auth/internal/auth/iam"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpHandlerPrincipalMissing = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_handler_principal_missing_total",
		Help: "IAM HTTP handler could not find any principal token",
	})
	httpHandlerExtractionSkipped = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_handler_extraction_skipped_total",
		Help: "IAM HTTP handler ignored missing principal token",
	})
	httpHandlerExtractionFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_handler_extraction_failed_total",
		Help: "IAM HTTP handler failed to extract principal token",
	})
	httpHandlerPrincipalExtracted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_handler_extracted_principal_total",
		Help: "IAM HTTP handler successfully extracted principal",
	})
)

// HTTPHandler propagates the token from the HTTP authorization header into the context.
//
// `must` is used to indicate that the token must be present.
func HTTPHandler(must bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("authorization")
			if header == "" {
				if must {
					httpHandlerPrincipalMissing.Inc()
					writeErr(w, httpError{
						Code:    http.StatusUnauthorized,
						Message: "missing authorisation token",
					})
					return
				}

				// allow unauthenticated requests in
				httpHandlerExtractionSkipped.Inc()
				next.ServeHTTP(w, r)
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) < 2 {
				httpHandlerExtractionFailed.Inc()
				writeErr(w, httpError{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				})
				return
			}

			ctx := iam.ToIncomingCtx(r.Context(), parts[1])
			ctx = iam.ToOutgoingCtx(ctx, parts[1])
			r = r.WithContext(ctx)
			httpHandlerPrincipalExtracted.Inc()

			next.ServeHTTP(w, r)
		})
	}
}

type httpError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func writeErr(w http.ResponseWriter, e httpError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		fmt.Printf("failed to write error: %s\n", err)
	}
}
