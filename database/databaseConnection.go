package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	//load .env file
	err := godotenv.Load(".env")
	//check for error while loading .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//get url
	MongoDbURL := os.Getenv("MONGODB_URL")

	//create a new client to connect to a deployment specified by the uri
	//options.Client().ApplyURI allow the application of mongodburl
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(MongoDbURL))

	if err != nil {
		log.Fatal(err)
	}

	//create context with timeout set
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	//Ensure connnection is closed at the end of processing
	defer cancelFunc()

	//initializes the mongo Client
	err = mongoClient.Connect(ctx)
	//check for error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return mongoClient

}

//globally available ref to the mongo client instance
var Client *mongo.Client = DBinstance()

//return a mongo collection..
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	//handler for a database with the given name configured with the given DatabaseOptions.

	//uses the handler to create a collection
	collection := client.Database("restaurantDB").Collection(collectionName)

	return collection
}
