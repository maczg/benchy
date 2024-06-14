package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/massimo-gollo/benchy/services/simple"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	readTimeout  int
	writeTimeout int
)

var simpleCmd = cobra.Command{
	Use:   "simple",
	Short: "Start a simple http server",
	Long:  `Start a simple http server. Expose two endpoints: /cpu and /memory to test CPU and memory intensive tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.New()
		r.Use(gin.Recovery())
		r.Use(loggerWithFormat)
		r.Use(middleware.LoggerGin())
		r.GET("/healthz", simple.HealthHandler)
		r.GET("/cpu", simple.CpuTaskHandler)
		r.GET("/cpuintensive", simple.CpuIntensiveTaskHandler)
		r.GET("/mem", simple.MemTaskHandler)

		s := &http.Server{
			Addr:           ":8080",
			Handler:        r,
			ReadTimeout:    time.Duration(readTimeout) * time.Second,
			WriteTimeout:   time.Duration(writeTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		// Channel to listen for interrupt or terminate signals
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// Run server in a goroutine so that it doesn't block
		go func() {
			if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logrus.Fatalf("Could not listen on %s: %v\n", ":8080", err)
			}
		}()
		logrus.Println("Server is ready to handle requests at :8080")

		// Block until we receive a signal
		<-stop
		logrus.Println("Shutting down server...")

		// Create a deadline to wait for server to shut down gracefully
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown the server gracefully
		if err := s.Shutdown(ctx); err != nil {
			logrus.Fatalf("Server forced to shutdown: %v", err)
		}

		logrus.Println("Server exiting")
	},
}

func init() {
	simpleCmd.AddCommand(version.Command())
	simpleCmd.PersistentFlags().StringVarP(&serverPort, "port", "p", "8080", "simple server port to listen on")
	simpleCmd.PersistentFlags().IntVarP(&readTimeout, "read-timeout", "r", 10, "simple server read timeout in second")
	simpleCmd.PersistentFlags().IntVarP(&writeTimeout, "write-timeout", "w", 100, "simple server write timeout in second")
}

var loggerWithFormat = gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] %s %d %s %s \n",
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.StatusCode,
		param.Latency,
		param.ErrorMessage,
	)
})
