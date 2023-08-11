package middleware

import (
	"context"
	uid "github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CtxKeyRequestID struct{}
type CtxKeyLog struct{}

func LoggerMw(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return loggerMw(next, logger)
	}
}

func loggerMw(next http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var requestID string
		if requestID = r.Header.Get("X-Request-ID"); requestID == "" {
			uuid, _ := uid.NewRandom()
			requestID = uuid.String()
			ctx = context.WithValue(ctx, CtxKeyRequestID{}, requestID)
			w.Header().Set("X-Request-ID", requestID)
		}

		start := time.Now()

		rr := &responseRecorder{w: w}

		l := logger.WithFields(logrus.Fields{
			"req.path":   r.URL.Path,
			"req.method": r.Method,
			"req.id":     requestID,
		})

		l.Debugln("request started")

		defer func() {
			l.WithFields(logrus.Fields{
				"resp.took_ms": int64(time.Since(start) / time.Millisecond),
				"resp.status":  rr.status,
				"resp.bytes":   rr.b}).Infoln("request complete")
		}()

		ctx = context.WithValue(ctx, CtxKeyLog{}, l)
		r = r.WithContext(ctx)
		next.ServeHTTP(rr, r)
	})
}
