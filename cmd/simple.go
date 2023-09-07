package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/massimo-gollo/benchy/services/simple"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
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

		logrus.Fatalln(s.ListenAndServe())
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
