package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getTrainersCollection() *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("trainers")

	return collection
}

var trainers = getTrainersCollection()

func pong(ctx iris.Context) {
	board := ctx.Params().Get("board")
	thread := ctx.Params().Get("thread")
	post := ctx.Params().Get("post")

	ctx2, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := trainers.InsertOne(ctx2, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID

	ctx.JSON(iris.Map{
		"id":    id,
		"board": board,
		"thread": thread,
		"post":  post,
	})
}

func main() {
	app := iris.Default()
	// Resource Management
	app.Delete("/{board:string}/{thread:string}/{post:string}", pong) // Delete a resource
	app.Delete("/{board:string}/{thread:string}", pong)               // Delete a thread
	app.Post("/{board:string}/{thread:string}", pong)                 // Post to a thread
	app.Get("/{board:string}/{thread:string}", pong)                  // Get all post of a thread

	// Thread Management
	app.Get("/{board:string}", pong)  // Get all threads of a Board
	app.Post("/{board:string}", pong) // Post a thread to a Board

	// Board Management
	app.Delete("/{board:string}", pong) // Delete a board
	app.Get("/", pong)                  // Get all Boards
	app.Post("/", pong)                 // Create a Board

	app.Run(iris.Addr(":8080"))
}
