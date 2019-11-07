package boards

import (
	"context"
	"fmt"
	"log"

	"github.com/pedroorez/godump/mehchan/mongo"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

// Board is the basic board structure
type Board struct {
	Name string
}

// GetBoards get all boards
func GetBoards(ctx iris.Context) {
	var results []*Board

	collection := mongo.GetDatabaseClient().Collection("board")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Board
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(iris.Map{"chave": "valor"})
}

// PostBoard creates a board.
func PostBoard(ctx iris.Context) {
	var board Board
	if err := ctx.ReadJSON(&board); err != nil {
		log.Fatal(err)
	}
	collection := mongo.GetDatabaseClient().Collection("board")
	insertResult, err := collection.InsertOne(context.TODO(), board)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	ctx.JSON(iris.Map{"id": insertResult.InsertedID, "board": board})
}
