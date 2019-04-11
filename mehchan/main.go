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

	ctx2, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := trainers.InsertOne(ctx2, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID

	ctx.JSON(iris.Map{
		"id": id,
	})
}

func main() {
	app := iris.Default()
	app.Post("/{board}/{thread}", pong)
	app.Get("/{board}/{thread}", pong)

	app.Get("/{board}", pong)
	app.Post("/{board}", pong)

	app.Run(iris.Addr(":8080"))
}
