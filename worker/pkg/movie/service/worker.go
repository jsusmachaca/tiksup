package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/eventstream/model"
	"github.com/jsusmachaca/tiksup/pkg/movie/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MovieWorker(db *sql.DB, kafkaData model.KafkaData, mC model.MongoConnection) {
	collection := mC.Collection
	ctx := mC.CTX

	movie := repository.MovieRository{DB: db}
	user_id := kafkaData.UserID

	recomendation, err := movie.GetPreferences(user_id)
	if err != nil {
		log.Panicln("Error to obtain recomendations", err)
	}
	history, err := movie.GetHistory(user_id)
	if err != nil {
		log.Println("No history", err)
	}

	filter := bson.M{"_id": bson.M{"$nin": history}}
	findOptions := options.Find()
	findOptions.SetLimit(6)

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &recomendation.Movies)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(recomendation, "", "    ")
	if err != nil {
		log.Panicln("Error to marshall recomendations", err)
	}

	fmt.Printf("%s\n", jsonData)
}
