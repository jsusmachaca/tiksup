package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/eventstream/model"
	"github.com/jsusmachaca/tiksup/pkg/movie/repository"
)

func MovieWorker(db *sql.DB, kafkaData model.KafkaData, mC model.MongoConnection) {
	movie := repository.MovieRository{DB: db}
	mongoMovie := repository.MongoRepository{Collection: mC.Collection, CTX: mC.CTX}

	user_id := kafkaData.UserID
	recomendation, err := movie.GetPreferences(user_id)
	if err != nil {
		log.Fatal(err)
	}
	history, err := movie.GetHistory(user_id)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoMovie.GetMoviesExcludeHistory(history, &recomendation.Movies)
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(recomendation)
	if err != nil {
		log.Fatal(err)
	}
	bodyReader := bytes.NewReader(body)

	err = ApiService(bodyReader)
	if err != nil {
		log.Fatal(err)
	}
}
