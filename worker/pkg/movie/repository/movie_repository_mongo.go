package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
	CTX        context.Context
}

func (movie *MongoRepository) GetMoviesExcludeHistory(history []primitive.ObjectID, movies any) error {
	filter := bson.M{"_id": bson.M{"$nin": history}}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 6},
		}}},
	}

	cursor, err := movie.Collection.Aggregate(movie.CTX, pipeline)
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

func (movie *MongoRepository) GetRadomMovies(movies any) error {
	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 6},
		}}},
	}

	cursor, err := movie.Collection.Aggregate(movie.CTX, pipeline)
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
