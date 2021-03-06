package main

import (
	"fmt"

	"github.com/dfrezell/patzer"
)

func main() {
	patzer.Cfg.Parse()

	board := patzer.NewBoard()
	//board := patzer.EmptyBoard()
	//board = board.Place(patzer.Piece{patzer.WHITE, patzer.KING}, patzer.E1)
	//board = board.Place(patzer.Piece{patzer.BLACK, patzer.KING}, patzer.E8)
	fmt.Printf("FEN = %s\n", *board.FEN())
	fmt.Printf("%s\n", *board.ASCII())
	board, _ = board.Move(patzer.E2, patzer.E4)
	fmt.Printf("FEN = %s\n", *board.FEN())
	fmt.Printf("%s\n", *board.ASCII())
	board, _ = board.Move(patzer.E7, patzer.E5)
	fmt.Printf("FEN = %s\n", *board.FEN())
	fmt.Printf("%s\n", *board.ASCII())
	board, _ = board.Move(patzer.G1, patzer.F3)
	fmt.Printf("FEN = %s\n", *board.FEN())
	fmt.Printf("%s\n", *board.ASCII())
	board, _ = board.Move(patzer.B8, patzer.C6)
	fmt.Printf("FEN = %s\n", *board.FEN())
	fmt.Printf("%s\n", *board.ASCII())

	board.GenMoves()
}
