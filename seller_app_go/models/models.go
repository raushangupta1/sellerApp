package models

import (
	"context"
	"log"

	"github.com/gobuffalo/envy"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var MongoDB *mongo.Database

func init() {
	var err error
	url := envy.Get("MONGO_URL", "mongodb")
	dbName := envy.Get("MONGO_DATABASE_NAME", "local")
	clientOptions := options.Client().ApplyURI(url)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	MongoDB = client.Database(dbName)
}

func GetClient() *mongo.Client {
	return client
}
