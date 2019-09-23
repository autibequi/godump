package main

import (
	"github.com/kataras/iris"
	"github.com/pedroorez/godump/mehchan/boards"
)

func main() {
	app := iris.Default()

	app.Get("/boards", boards.GetBoards)
	app.Post("/boards", boards.PostBoard)

	app.Run(iris.Addr(":8080"))
}
