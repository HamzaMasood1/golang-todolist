package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	const uri = "mongodb://localhost:27017"
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
}
func listAll() {
	fmt.Println("hello")
}
func main() {
	//greetings
	fmt.Println("Welcome to your todolist app\nPlease pick from the following options")
	fmt.Println("1. List all todolists")
	fmt.Println("2. Create a todolist")
	fmt.Println("3. Edit a todolist")
	fmt.Println("4. Delete a todolist")

	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		listAll()
	}
}
