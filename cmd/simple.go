package cmd

import (
	"github.com/massimo-gollo/benchy/internal/simple"
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/spf13/cobra"
	"strconv"
)

var simplePort int

var simpleCmd = cobra.Command{
	Use:   "simple",
	Short: "Start a simple http server",
	Long:  `Start a simple http server. Expose two endpoints: /cpu and /memory to test CPU and memory intensive tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := simple.NewServer(simple.ServiceName, strconv.Itoa(simplePort))
		s.AddHandler("/cpu", simple.CpuIntensiveHandler)
		s.AddHandler("/memory", simple.MemoryIntensiveHandler)
		s.AddMiddleware(middleware.LoggerMw(s.Logger()))
		s.Start()
	},
}

func init() {
	simpleCmd.PersistentFlags().IntVarP(&simplePort, "port", "p", 8080, "port to listen on")
}
