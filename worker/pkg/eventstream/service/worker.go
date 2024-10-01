package service

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jsusmachaca/tiksup/internal/config"
	modelKafka "github.com/jsusmachaca/tiksup/pkg/eventstream/model"
	"github.com/jsusmachaca/tiksup/pkg/eventstream/repository"
	movieService "github.com/jsusmachaca/tiksup/pkg/movie/service"
)

func KafkaWorker(configMap *kafka.ConfigMap, db *sql.DB, mC modelKafka.MongoConnection) error {
	var kafkaData modelKafka.KafkaData
	kafaDB := repository.KafkaRepository{DB: db}

	consumer, err := config.KafKaConsumer(configMap)
	if err != nil {
		log.Fatalf("Kafka worker error: %v", err)
	}

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error to trying get kafka information: %v\n", err)
		}
		json.Unmarshal(msg.Value, &kafkaData)
		log.Printf("message received: %s\n", msg.Value)

		if err := kafaDB.UpdateUserInfo(kafkaData); err != nil {
			log.Printf("Error to insert kafka information: %v\n", err)
		}
		if kafkaData.Next {
			go movieService.MovieWorker(db, kafkaData, mC)
		}
	}
}
