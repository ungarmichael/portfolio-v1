package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var clientOptions *clientOptions
var ungarClient *mongo.Client
var ctx = context.TODO()

func ConnectToMongoDB() {
	var err error
	var ungarClient *mongo.Client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ungarClient, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://ungar:ungardatabase_123@cluster0.ykrvq.mongodb.net/<Ungar>?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// ungarClient, err = mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = ungarClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("successfully connected")
}

func DisconnectFromMongoDB() {
	err := ungarClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
