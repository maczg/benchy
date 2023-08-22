package simple

import (
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/spf13/cobra"
	"strconv"
)

var simplePort int

var SimpleCmd = cobra.Command{
	Use:   "simple",
	Short: "Start a simple http server",
	Long:  `Start a simple http server. Expose two endpoints: /cpu and /memory to test CPU and memory intensive tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := NewServer(ServiceName, strconv.Itoa(simplePort))
		s.AddHandler("/cpu", CpuIntensiveHandler)
		s.AddHandler("/memory", MemoryIntensiveHandler)
		s.AddMiddleware(middleware.LoggerMw(s.Logger()))
		s.Start()
	},
}

func init() {
	SimpleCmd.PersistentFlags().IntVarP(&simplePort, "port", "p", 8080, "port to listen on")
	SimpleCmd.AddCommand(version.Command())
}
