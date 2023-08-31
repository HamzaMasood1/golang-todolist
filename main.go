package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

type Todolist struct {
	Name     string      `bson:"name"`
	Todolist interface{} `bson:"todolist"`
}

var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

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
	cursor.All(context.TODO(), &results)
	var t Todolist
	fmt.Println("\nPrinting all todolists:")
	for r, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &t)
		fmt.Println(result)
		fmt.Println(t)
		fmt.Printf("%v %v\n", r+1, t.Name)
	}
	fmt.Println("")
}

func createtodolist() {
	fmt.Println("What would you like to name your todolist?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Please input the entries of your todolist")
	cont := true
	var input string
	entries := make([]string, 0)
	count := 1
	for cont {
		fmt.Print(count, "\t")
		reader := bufio.NewReader(os.Stdin)
		sentence, _ := reader.ReadString('\n')
		entries = append(entries, sentence)
		fmt.Println("Would you like to add more entries? (yes/y)")
		fmt.Scanln(&input)
		if strings.EqualFold(input, "yes") || strings.EqualFold(input, "y") {
			count++
			continue
		} else {
			cont = false
		}
	}
	fmt.Println("Name of todolist:", name)
	fmt.Println("Entries of your todolist:")
	fmt.Println(entries)

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
	case 2:
		createtodolist()
	}
}
