package main

import (
	"github.com/pedroorez/godump/mehchan/boards"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/boards", boards.GetBoards)
	app.Post("/boards", boards.PostBoard)

	app.Run(iris.Addr(":8080"))
}
