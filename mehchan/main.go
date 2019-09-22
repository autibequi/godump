package main

import (
	"github.com/kataras/iris"

	Boards "github.com/pedroorez/godump/mehchan/boards"
)

func main() {
	app := iris.Default()

	app.Get("/boards", Boards.GetBoards)
	app.Post("/boards", Boards.PostBoard)

	app.Run(iris.Addr(":8080"))
}
