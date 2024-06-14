package kafka

import (
	"context"
	"crypto/tls"
	"github.com/Shopify/sarama"
	zlog "log"
	"os"
)

var (
	Topic           string
	ProtocolVersion = sarama.V3_0_0_0
	GroupID         string

	log = zlog.New(os.Stdout, "kafka-consumer", zlog.LstdFlags)
)

func StartConsumerGroup(ctx context.Context, brokers []string) error {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = ProtocolVersion
	// So we can know the partition and offset of messages.
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Net.TLS.Enable = true
	saramaConfig.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, GroupID, saramaConfig)
	if err != nil {
		return err
	}

	c := Consumer{
		log: log,
	}

	err = consumerGroup.Consume(ctx, []string{Topic}, &c)
	if err != nil {
		return err
	}
	return nil
}

type Consumer struct {
	log *zlog.Logger
}

func (g *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	g.log.Println("Consumer setup")
	return nil
}

func (g *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	g.log.Println("Consumer cleanup")

	return nil
}

func (g *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Println("message channel was closed")
				return nil
			}

			log.Printf("Message claimed: orderId = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
