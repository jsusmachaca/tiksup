package config

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConfig() kafka.ConfigMap {
	server := os.Getenv("KAFKA_SERVER")
	topic := os.Getenv("KAFKA_TOPIC")
	fmt.Printf("Kafka Server: %s\n", server)
	fmt.Printf("Kafka Topic: %s\n", topic)

	configMap := kafka.ConfigMap{
		"bootstrap.servers":  server,
		"group.id":           "user-info-consumer",
		"auto.offset.reset":  "latest",
		"session.timeout.ms": 300000,
	}
	return configMap
}

func KafKaConsumer(configMap *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	topic := os.Getenv("KAFKA_TOPIC")
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	log.Println("\033[32mSUBSCRIBE TO KAFKA TOPIC\033[0m")

	return consumer, nil
}
