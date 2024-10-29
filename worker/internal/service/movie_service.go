package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jsusmachaca/tiksup/pkg/eventstream"
	"github.com/jsusmachaca/tiksup/pkg/movie"
)

func MovieWorker(client *http.Client, db *sql.DB, kafkaData eventstream.KafkaData, mongoConn movie.MongoConnection) {
	movieRepository := &movie.MovieRepository{DB: db}
	mongoMovie := mongoConn.ToRepository()

	user_id := kafkaData.UserID
	recomendation, err := movieRepository.GetPreferences(user_id)
	if err != nil {
		log.Fatal(err)
	}
	history, err := movieRepository.GetHistory(user_id)
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

	err = movie.ApiService(client, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("\033[32mmovies send to data processor\033[0m")
}
