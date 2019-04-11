package main

import "github.com/kataras/iris"

func pong(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "pong",
	})
}

func main() {
	app := iris.Default()
	app.Get("/ping", pong)
	app.Run(iris.Addr(":8080"))
}
