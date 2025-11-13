package logger

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// forensicLog context log type
type forensicLog string

// forensicLogKey key for context storage
var forensicLogKey = forensicLog("flog")

// ForensicLoggerFromRequest provides a forensic logger from a request
func ForensicLoggerFromRequest(r *http.Request) *log.Logger {
	return ForensicLogFromContext(r.Context())
}

// ForensicLogFromContext provides a forensic log from a context
func ForensicLogFromContext(ctx context.Context) *log.Logger {
	logger := ctx.Value(forensicLogKey)
	if logger == nil {
		logger = log.StandardLogger()
	}

	return logger.(*log.Logger)
}
