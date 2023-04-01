package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func NewLoggingMiddleware(logger logrus.FieldLogger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&logFormatter{logger: logger})
}

type logFormatter struct {
	logger logrus.FieldLogger
}

func (l *logFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	logFields := logrus.Fields{
		"http_method": r.Method,
		"remote_addr": r.RemoteAddr,
		"uri":         r.RequestURI,
	}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	entry := &logEntry{
		logger: l.logger.WithFields(logFields),
	}

	entry.logger.Info("Request started")

	return entry
}

type logEntry struct {
	logger logrus.FieldLogger
}

func (l *logEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	logger := l.logger.WithFields(logrus.Fields{
		"resp_status":       status,
		"resp_bytes_length": bytes,
		"resp_elapsed":      elapsed.Round(time.Millisecond / 100).String(),
	})

	logger.Info("Request completed")
}

func (l *logEntry) Panic(v interface{}, stack []byte) {
	l.logger = l.logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}
