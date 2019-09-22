package Boards

import (
	"github.com/kataras/iris"
)

type Board struct {
	Name string
}

func GetBoards(ctx iris.Context) {
	board := Board{Name: "Teste"}
	ctx.JSON(board)
}

func PostBoard(ctx iris.Context) {
	board := Board{Name: "Teste"}
	ctx.JSON(board)
}
