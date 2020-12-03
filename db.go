package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbURI the URI of the database server
const DbURI = "mongodb://mongo:27017"

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI(DbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
