package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConnection struct {
	Collection *mongo.Collection
	CTX        context.Context
}
