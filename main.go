package main

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

type todolist struct {
	name     string
	todolist []string
}

func init() {
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
func printNumberofToDoLists() int64 {
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("mydb").Collection("todolist")
	opts2 := options.Count().SetHint("_id_")
	count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts2)
	if err != nil {
		panic(err)
	}
	return count
}
func listAll() {
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("mydb").Collection("todolist")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var results = make([]interface{}, 0)
	type Todo struct {
		ID       primitive.ObjectID `bson:"_id"`
		Name     interface{}        `bson:"name"`
		Todolist interface{}        `bson:"todolist"`
	}
	cursor.All(context.TODO(), &results)
	var t Todo
	for _, result := range results {
		fmt.Println(reflect.TypeOf(result))
		fmt.Println(result)
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &t)
		fmt.Println(t)

	}
}

func main() {

	fmt.Println("Welcome to your todolist app")
	fmt.Println("Number of todo lists: ", printNumberofToDoLists())
	fmt.Println("Please pick from the following options")
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
