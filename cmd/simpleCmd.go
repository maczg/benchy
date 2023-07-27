package cmd

import (
	"github.com/massimo-gollo/benchy/server"
	"github.com/massimo-gollo/benchy/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool

var simpleCmd = cobra.Command{
	Use:   "simple",
	Short: "Start a simple benchmark server",
	Long:  `Start a simple benchmark server`,
	Run: func(cmd *cobra.Command, args []string) {
		port := utils.GetEnvOrDefault("PORT", "3000")
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		s := server.New(port)
		s.AddHandler("/cpu", s.CpuIntensiveHandler)
		s.AddHandler("/memory", s.MemoryIntensiveHandler)
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(&simpleCmd)
	simpleCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

}
