package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Collection *mongo.Collection
	CTX        context.Context
}

func (movie *MongoRepository) GetMoviesExcludeHistory(history []primitive.ObjectID, movies any) error {
	filter := bson.M{"_id": bson.M{"$nin": history}}
	findOptions := options.Find()
	findOptions.SetLimit(6)

	cursor, err := movie.Collection.Find(movie.CTX, filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(movie.CTX)

	err = cursor.All(movie.CTX, movies)
	if err != nil {
		return err
	}
	return nil
}
