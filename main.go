package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string `bson:"name"`
}

func main() {
	// Set client options for mongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//* Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("easymongo").Collection("human")

	//!------------------------CRUD in go mongo driver------------------------------

	//! Insert a document
	person := Person{Name: "John"}
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//! Update a document
	filter := bson.M{"name": "John"}
	update := bson.M{"$set": bson.M{"name": "Jane"}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//! Find a document
	var result Person
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	//! Delete a document
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents.\n", deleteResult.DeletedCount)

	//!----------------------------------------------------------------------------------

	//* Disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
