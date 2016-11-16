package main

import (
	"fmt"

	"github.com/dfrezell/patzer"
)

func main() {
	patzer.Cfg.Parse()

	board := patzer.NewBoard()
	fmt.Printf("FEN = %s\n", *board.FEN())

	board.GenMoves()
}
