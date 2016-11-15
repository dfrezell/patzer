package main

import (
    "fmt"

    "github.com/dfrezell/patzer"
)

func main() {
    patzer.Cfg.Parse()

    board := patzer.NewBoard()
    fmt.Printf("FEN = %s\n", *board.FEN())

	nb, err := board.Move(patzer.E2, patzer.E4)
	if err == nil {
		fmt.Printf("b = %p\nn = %p\n", board, nb)
	}
}
