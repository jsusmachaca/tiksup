package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/eventstream/model"
	"github.com/jsusmachaca/tiksup/pkg/movie/repository"
)

func MovieWorker(db *sql.DB, kafkaData model.KafkaData) {
	movie := repository.MovieRository{DB: db}
	user_id := kafkaData.UserID

	recomendation, err := movie.GetPreferences(user_id)
	if err != nil {
		log.Panicln("Error to obtain recomendations", err)
	}

	jsonData, err := json.MarshalIndent(recomendation, "", "    ")
	if err != nil {
		log.Panicln("Error to marshall recomendations", err)
	}

	history, err := movie.GetHistory(user_id)
	if err != nil {
		log.Println("No history", err)
	}

	fmt.Printf("%s\n", jsonData)
	fmt.Printf("%s\n", history)
}
