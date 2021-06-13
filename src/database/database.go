package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConn mongo.Client

var mongoClient *mongoConn

func GetMongoDBConn() *mongoConn {
	return mongoClient
}

func (conn *mongoConn) Client() *mongo.Client {
	client := mongo.Client(*conn)
	return &client
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

	mongoConnFromClient := mongoConn(*client)
	mongoClient = &mongoConnFromClient
}
