package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func GetMongoDBClient() *mongo.Client {
	return mongoClient
}

func Connect() {
	connectionString := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(connectionString).SetMaxPoolSize(50)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client
}
