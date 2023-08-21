package cmd

import (
	"github.com/massimo-gollo/benchy/internal/standalone"
	"github.com/massimo-gollo/benchy/pkg/middleware"
	"github.com/spf13/cobra"
	"strconv"
)

var standAlonePort int

var standaloneCmd = cobra.Command{
	Use:   "standalone",
	Short: "Start a simple standalone http server",
	Long:  `Start a simple standalone http server. Expose two endpoints: /cpu and /memory to test CPU and memory intensive tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		//port := mixin.GetEnvOrDefault("PORT", "3000")
		s := standalone.New(standalone.ServiceName, strconv.Itoa(standAlonePort))
		s.AddHandler("/cpu", standalone.CpuIntensiveHandler)
		s.AddHandler("/memory", standalone.MemoryIntensiveHandler)
		s.AddMiddleware(middleware.LoggerMw(s.Logger()))
		s.Start()
	},
}

func init() {
	standaloneCmd.PersistentFlags().IntVarP(&standAlonePort, "port", "p", 8080, "port to listen on")
}
