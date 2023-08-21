package trace

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func Middleware(tracer trace.Tracer, logger *logrus.Entry) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debugln("tracing middleware called")

			ctx := r.Context()

			span := trace.SpanFromContext(ctx)
			if span.SpanContext().IsValid() {
				logger.Infoln("Existing root span found.")
			} else {
				// No existing span found, so create a new root span
				url := r.URL.String()
				ctx, span = tracer.Start(ctx, url)
				logger.Infoln("New root span created.")
			}
			defer span.End()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
