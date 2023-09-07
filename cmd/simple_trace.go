package cmd

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/massimo-gollo/benchy/pkg/log"
	"github.com/massimo-gollo/benchy/pkg/tracing"
	"github.com/massimo-gollo/benchy/pkg/util"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/massimo-gollo/benchy/services/simpletrace"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

var traceCmd = cobra.Command{
	Use:   "simpletrace",
	Short: "Start simple server with tracing",
	Long:  "Start simple server instrumented with otel for tracing",
	Run: func(cmd *cobra.Command, args []string) {

		util.MustMapEnv(&simpletrace.OtelEndpoint, "OTEL_EXPORTER_OTLP_ENDPOINT")

		// alias for logger
		logger := simpletrace.Logger

		logger.Infoln("starting server: init tracing")

		shutdown, err := tracing.InitOLTPTracer(simpletrace.ServiceName, simpletrace.OtelEndpoint)
		if err != nil {
			logger.WithError(err).Fatalln("failed to initialize tracer")
		}
		defer func() {
			if err := shutdown(context.Background()); err != nil {
				logger.WithError(err).Errorln("failed to shutdown tracer")
			}
		}()

		tracer := otel.Tracer(simpletrace.ServiceName)

		r := mux.NewRouter()
		r.Use(tracing.Middleware(tracer, log.SetLogName(simpletrace.ServiceName, logger)))

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

		logger.Infof("starting server: listening on port %s", serverPort)
		logger.Fatalln(http.ListenAndServe(":"+serverPort, r))

	},
}

func init() {
	traceCmd.AddCommand(version.Command())
	traceCmd.PersistentFlags().StringVarP(&simpletrace.ServiceName, "name", "n", "simpletrace", "service name")
	traceCmd.PersistentFlags().StringVarP(&serverPort, "port", "p", "8080", "simple server port to listen on")
}

func doWork(ctx context.Context, tracer trace.Tracer) {
	ctx, span := tracer.Start(
		ctx,
		"work-test-handler-1")
	defer span.End()
	// emulate doing work
	time.Sleep(200 * time.Millisecond)
}
