package main

import (
	"context"
	"time"
        // "labix.org/v2/mgo/bson"
	"github.com/kataras/iris"
        "github.com/pedroorez/godump/mongo"
)

var client = mongo.GetMongoClient()
var collection = client.Database("testing").Collection("numbers")

func add(irisctx iris.Context) {
	texto := irisctx.Params().Get("texto")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": texto})
	id := res.InsertedID
	irisctx.JSON(iris.Map{"message": id})
}

func list(ctx iris.Context) {
	texto := ctx.Params().Get("texto")
	ctx.JSON(iris.Map{"message": texto})
}


func main() {
	app := iris.Default()
	app.Get("/add/{texto:string}", add)
	app.Get("/list/", list)
	app.Run(iris.Addr(":8080"))
}
