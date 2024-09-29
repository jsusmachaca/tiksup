package service

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jsusmachaca/tiksup/internal/config"
	modelKafka "github.com/jsusmachaca/tiksup/pkg/eventstream/model"
)

func KafkaWorker(configMap *kafka.ConfigMap) error {
	var kafkaData modelKafka.KafkaData

	consumer, err := config.KafKaConsumer(configMap)
	if err != nil {
		log.Fatalf("Kafka worker error: %v", err)
	}

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		json.Unmarshal(msg.Value, &kafkaData)
		log.Printf("message received: %s\n", msg.Value)
	}
}
