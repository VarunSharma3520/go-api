package db

import (
	"github.com/VarunSharma3520/go-api/internal/config"
	// "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
var Client *mongo.Client

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.GetConfig().MongoURI)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
