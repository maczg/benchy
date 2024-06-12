package cmd

import (
	"github.com/massimo-gollo/benchy/pkg/util"
	"github.com/massimo-gollo/benchy/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

var (
	serverPort string
	redisAddr  string
	otelAddr   string

	kafkaAddr  string
	kafkaTopic string
	kafkaGroup string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "benchy",
	Short: "Benchy, CLI benchmark launcher",
	Long:  `Benchy is a CLI launcher to start different kind of benchmark microservices.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(&simpleCmd)
	rootCmd.AddCommand(&traceCmd)
	rootCmd.AddCommand(&kafkaConsumerCmd)
	rootCmd.AddCommand(&loadCmd)
	rootCmd.AddCommand(version.Command())

	initVariables()

}

func initVariables() {
	serverPort = util.GetEnvOrDefault("SERVER_PORT", "8080")
	redisAddr = util.GetEnvOrWarning("REDIS_ADDR")
	otelAddr = util.GetEnvOrWarning("OTEL_ADDR")
	kafkaAddr = util.GetEnvOrWarning("KAFKA_ADDR")
	kafkaTopic = util.GetEnvOrWarning("KAFKA_TOPIC")
	kafkaGroup = util.GetEnvOrWarning("KAFKA_GROUP")
}
