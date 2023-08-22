package srvtrace

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/massimo-gollo/benchy/pkg/log"
	"github.com/massimo-gollo/benchy/pkg/mixin"
	"github.com/massimo-gollo/benchy/pkg/trace"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	tr "go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

var TraceSrvCmd = cobra.Command{
	Use:   "server",
	Short: "Start simple server with tracing",
	Long:  "Start simple server instrumented with otel for tracing",
	Run: func(cmd *cobra.Command, args []string) {

		logger := log.NewDefaultLogger()
		logger.Infoln("starting server: init tracing")

		var otlpEndpoint = mixin.GetEnvOrDefault("OTLP_ENDPOINT", "localhost:4317")

		ctx := context.Background()

		shutdown, err := trace.InitOLTPTracer("server", otlpEndpoint)
		if err != nil {
			logger.WithError(err).Fatalln("failed to initialize tracer")
		}
		defer func() {
			if err := shutdown(ctx); err != nil {
				logger.WithError(err).Errorln("failed to shutdown tracer")
			}
		}()

		tracer := otel.Tracer("server")

		r := mux.NewRouter()
		r.Use(trace.Middleware(tracer, log.SetLogName("tracer", logger)))

		r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok\n"))
		})

		r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, span := tracer.Start(
				ctx,
				"work-test-handler")
			defer span.End()

			doWork(ctx, tracer)

			// emulate latency work
			time.Sleep(100 * time.Millisecond)

			_, _ = w.Write([]byte("test-handler\n"))
		})

		logger.Infoln("starting server: listening on :8080")
		logger.Fatalln(http.ListenAndServe(":8080", r))

	},
}

func init() {
	TraceSrvCmd.AddCommand(version.Command())
}

func doWork(ctx context.Context, tracer tr.Tracer) {
	ctx, span := tracer.Start(
		ctx,
		"work-test-handler-1")
	defer span.End()
	// emulate doing work
	time.Sleep(200 * time.Millisecond)
}
