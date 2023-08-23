package cmd

import (
	"context"
	"github.com/massimo-gollo/benchy/pkg/kafka"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
)

var kafkaConsumerCmd = cobra.Command{
	Use:   "kafka-consumer",
	Short: "Start kafka consumer",
	Long:  "Start kafka consumer",
	Run: func(cmd *cobra.Command, args []string) {

		logger := log.New(os.Stdout, "kafka-consumer", log.LstdFlags)

		logger.Println("starting kafka consumer")

		kafka.Topic = "otlp_spans"
		kafka.GroupID = "kafkaconsumer"
		brokers := []string{"kafka:9092"}

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		if err := kafka.StartConsumerGroup(ctx, brokers); err != nil {
			log.Fatal(err)
		}

		<-ctx.Done()

	},
}
