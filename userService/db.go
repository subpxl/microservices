package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func connectToDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	connString := os.Getenv("MONGO_URI")
	if connString == "" {
		log.Fatal("cannot find mongo db connection string")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()
	return client
}
