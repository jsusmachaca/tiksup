package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection(ctx context.Context) (*mongo.Collection, error) {
	MONGO_HOST := os.Getenv("MONGO_HOST")
	MONGO_PORT := os.Getenv("MONGO_PORT")
	MONGO_USER := os.Getenv("MONGO_USER")
	MONGO_PASSWORD := os.Getenv("MONGO_PASSWORD")
	MONGO_DB := os.Getenv("MONGO_DB")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)

	clientOpions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collection := client.Database(MONGO_DB).Collection("videos")

	return collection, nil
}
