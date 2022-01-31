package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoCN = ConnectToDataBase()
var clientOptions = getClientOptions()
const dbName string = "twittor"

func ConnectToDataBase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// Totalmente innecesario
	pingErr := client.Ping(context.TODO(), nil)

	if pingErr != nil {
		log.Fatal(pingErr.Error())
		return client
	}

	log.Println("Connected to DB")

	return client
}

func IsConnected() bool {
	pingErr := MongoCN.Ping(context.TODO(), nil)

	if pingErr != nil {
		return false
	}

	return true
}

func getClientOptions() *options.ClientOptions {
	connectionURI := os.Getenv("MONGO_URI")

	if connectionURI == "" {
		log.Fatal("Mongo connection uri was not found.")
		return options.Client()
	}

	return options.Client().ApplyURI(connectionURI)
}