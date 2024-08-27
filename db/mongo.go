package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoClient *mongo.Client

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}

func GetDB() *mongo.Database {
	client := GetMongoClient()
	return client.Database("tourney_share")
}
