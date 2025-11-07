package ctxlogger

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const idChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// forensicLog context log type
type forensicLog string

// forensicLogKey key for context storage
var forensicLogKey = forensicLog("flog")

type formatterWrapper struct {
	fid       string
	formatter log.Formatter
}

func (f *formatterWrapper) Format(e *log.Entry) ([]byte, error) {
	// We need add the `_fid` value in here, as the event is technically
	// immutable we make a data clone here so we don't mess things up
	data := make(log.Fields)
	for i, f := range e.Data {
		data[i] = f
	}

	data["_fid"] = f.fid

	entry := &log.Entry{
		Logger:  e.Logger,
		Time:    e.Time,
		Level:   e.Level,
		Message: e.Message,
		Buffer:  e.Buffer,
		Data:    data,
	}

	return f.formatter.Format(entry)
}

// NewContextWithForensicLog provides a forensic log attached to a context and the forensic ID
func NewContextWithForensicLog(ctx context.Context, idPrefix string) (context.Context, string) {

	b := make([]byte, 5)
	for i := range b {
		b[i] = idChars[rand.Intn(len(idChars))]
	}
	t := time.Now()

	fid := fmt.Sprintf(
		"%s%d%d%d%d%d%d%d%s",
		idPrefix,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		string(b),
	)

	logger := log.New()
	logger.Out = log.StandardLogger().Out
	logger.Level = log.StandardLogger().Level
	logger.Hooks = log.StandardLogger().Hooks
	logger.Formatter = &formatterWrapper{
		formatter: log.StandardLogger().Formatter,
		fid:       fid,
	}

	return context.WithValue(ctx, forensicLogKey, logger), fid
}

// ForensicLogFromContext provides a forensic log from a context
func ForensicLogFromContext(ctx context.Context) *log.Logger {
	logger := ctx.Value(forensicLogKey)
	if logger == nil {
		logger = log.StandardLogger()
	}

	return logger.(*log.Logger)
}

// NewForensicRequest provides a deep request clone with forensic logger added to the context
func NewForensicRequest(r *http.Request) (*http.Request, string) {
	ctx, fid := NewContextWithForensicLog(r.Context(), "req_")
	rc := r.WithContext(ctx)
	rc.Body = r.Body
	rc.Header = r.Header
	rc.Response = r.Response
	rc.Method = r.Method
	rc.Header.Add("X-ACS-Fid", fid)
	return rc, fid
}

// ForensicLoggerFromRequest provides a forensic logger from a request
func ForensicLoggerFromRequest(r *http.Request) *log.Logger {
	return ForensicLogFromContext(r.Context())
}
