package cmd

import (
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/massimo-gollo/benchy/services/simple"
	"github.com/spf13/cobra"
)

var simplePort string

var simpleCmd = cobra.Command{
	Use:   "simple",
	Short: "Start a simple http server",
	Long:  `Start a simple http server. Expose two endpoints: /cpu and /memory to test CPU and memory intensive tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := simple.NewServer(simple.ServiceName, simplePort)
		s.AddHandler("/cpu", simple.CpuIntensiveHandler)
		s.AddHandler("/memory", simple.MemoryIntensiveHandler)
		s.AddMiddleware(middleware.LoggerMw(s.Logger()))
		s.Start()
	},
}

func init() {
	simpleCmd.PersistentFlags().StringVarP(&simplePort, "port", "p", "8080", "simple server port to listen on")

}
