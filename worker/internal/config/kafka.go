package config

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConfig() kafka.ConfigMap {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVER"),
		"group.id":          "vote-consumer",
		"auto.offset.reset": "earliest",
	}
	return configMap
}

func KafKaConsumer(configMap *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe("tiksup-user-data", nil)
	if err != nil {
		return nil, err
	}
	log.Println("\033[32mSUBSCRIBE TO KAFKA TOPIC\033[0m")

	return consumer, nil
}
