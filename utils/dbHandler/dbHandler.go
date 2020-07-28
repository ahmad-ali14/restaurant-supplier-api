package dbHandler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	//"restaurant-supplier-api/config"
	"time"
)

func CreateConnection() *mongo.Client {
	// mongo conncetion
	//var mongoURI string = config.GetMongoUrl()
	var mongoURI string
	if os.Getenv("URI") != "" {
		mongoURI = os.Getenv("URI")
	} else {
		// mongoURI =	config.GetMongoUrl()

	}

	// if mongoURI != "" {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	//collection := client.Database("portfolio").Collection("restaurants")

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		//return "error, no connection"
	}

	// fmt.Printf("ctx is %s", ctx)

	return client
}
func AddItemToDb() string {
	return "Added"
}

func RemoveItemFromDb() string {
	return "Removed"
}
