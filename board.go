package patzer

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

const (
	NONE = 0
	NOSQ = MAX_SQUARES-1
)

const (
	WHITE = iota
	BLACK
	COLOR_MAX
)

const (
	_ = iota
	PAWN
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
	PIECES_MAX
)

var PieceSymbol [COLOR_MAX][PIECES_MAX]string = [COLOR_MAX][PIECES_MAX]string{
	[...]string{"", "P", "N", "B", "R", "Q", "K"},
	[...]string{"", "p", "n", "b", "r", "q", "k"},
}

const (
	A1 = iota
	    B1; C1; D1; E1; F1; G1; H1; _; _; _; _; _; _; _; _;
	A2; B2; C2; D2; E2; F2; G2; H2; _; _; _; _; _; _; _; _;
	A3; B3; C3; D3; E3; F3; G3; H3; _; _; _; _; _; _; _; _;
	A4; B4; C4; D4; E4; F4; G4; H4; _; _; _; _; _; _; _; _;
	A5; B5; C5; D5; E5; F5; G5; H5; _; _; _; _; _; _; _; _;
	A6; B6; C6; D6; E6; F6; G6; H6; _; _; _; _; _; _; _; _;
	A7; B7; C7; D7; E7; F7; G7; H7; _; _; _; _; _; _; _; _;
	A8; B8; C8; D8; E8; F8; G8; H8; _; _; _; _; _; _; _; _;
	MAX_SQUARES
)

var AlgebraicTable []string = []string{
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1", "-", "-", "-", "-", "-", "-", "-", "-",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2", "-", "-", "-", "-", "-", "-", "-", "-",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3", "-", "-", "-", "-", "-", "-", "-", "-",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4", "-", "-", "-", "-", "-", "-", "-", "-",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5", "-", "-", "-", "-", "-", "-", "-", "-",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6", "-", "-", "-", "-", "-", "-", "-", "-",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7", "-", "-", "-", "-", "-", "-", "-", "-",
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8", "-", "-", "-", "-", "-", "-", "-", "-",
}

const (
	RANK  = 16
	FILE  = 1
	LDIAG = RANK - FILE
	RDIAG = RANK + FILE
)

const (
	_      = iota
	RANK_1 = A1
	RANK_2 = A2
	RANK_3 = A3
	RANK_4 = A4
	RANK_5 = A5
	RANK_6 = A6
	RANK_7 = A7
	RANK_8 = A8
)

// mask use to determine if position is outside the bounds of the board
const OOB_MASK = 0x08

type Piece struct {
	Color int8
	Type  int8
}

var EmptyPiece Piece = Piece{NONE, NONE}

func validPosList(posList ...int) []int8 {
	// check if position is in an invalid square
	move := []int8{}
	for _, pos := range posList {
		if (pos < 0) || (pos > MAX_SQUARES) || ((pos & OOB_MASK) == OOB_MASK) {
			continue
		}
		move = append(move, int8(pos))
	}
	return move
}

func appendIf(moves [][]int8, vector []int8) [][]int8 {
	if len(vector) > 0 {
		moves = append(moves, vector)
	}
	return moves
}

func (p *Piece) Equal(other *Piece) bool {
	return p.Color == other.Color && p.Type == other.Type
}

func (p *Piece) whitePawnMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	if pos < RANK_3 {
		// two squares forward on 2nd rank
		endPos = appendIf(endPos, validPosList(x+RANK, x+RANK*2))
	} else {
		endPos = appendIf(endPos, validPosList(x+RANK))
	}
	endPos = appendIf(endPos, validPosList(x+LDIAG))
	endPos = appendIf(endPos, validPosList(x+RDIAG))
	return endPos
}

func (p *Piece) blackPawnMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	if pos >= RANK_7 {
		// two squares forward on 2nd rank
		endPos = appendIf(endPos, validPosList(x-RANK, x-RANK*2))
	} else {
		endPos = appendIf(endPos, validPosList(x-RANK))
	}
	endPos = appendIf(endPos, validPosList(x-LDIAG))
	endPos = appendIf(endPos, validPosList(x-RDIAG))
	return endPos
}

// we pretend a pawn can always capture on both diagonals
func (p *Piece) pawnMove(pos int8) [][]int8 {
	var endPos [][]int8
	if p.Color == WHITE {
		endPos = p.whitePawnMove(pos)
	} else if p.Color == BLACK {
		endPos = p.blackPawnMove(pos)
	} else {
		endPos = [][]int8{}
	}

	return endPos
}

func (p *Piece) knightMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	endPos = appendIf(endPos, validPosList(x-RDIAG-RANK))
	endPos = appendIf(endPos, validPosList(x-LDIAG-RANK))
	endPos = appendIf(endPos, validPosList(x-RDIAG-FILE))
	endPos = appendIf(endPos, validPosList(x-LDIAG+FILE))
	endPos = appendIf(endPos, validPosList(x+LDIAG-FILE))
	endPos = appendIf(endPos, validPosList(x+RDIAG+FILE))
	endPos = appendIf(endPos, validPosList(x+LDIAG+RANK))
	endPos = appendIf(endPos, validPosList(x+RDIAG+RANK))
	return endPos
}

func (p *Piece) bishopMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	endPos = appendIf(endPos, validPosList(x+LDIAG, x+LDIAG*2, x+LDIAG*3, x+LDIAG*4, x+LDIAG*5, x+LDIAG*6, x+LDIAG*7))
	endPos = appendIf(endPos, validPosList(x+RDIAG, x+RDIAG*2, x+RDIAG*3, x+RDIAG*4, x+RDIAG*5, x+RDIAG*6, x+RDIAG*7))
	endPos = appendIf(endPos, validPosList(x-LDIAG, x-LDIAG*2, x-LDIAG*3, x-LDIAG*4, x-LDIAG*5, x-LDIAG*6, x-LDIAG*7))
	endPos = appendIf(endPos, validPosList(x-RDIAG, x-RDIAG*2, x-RDIAG*3, x-RDIAG*4, x-RDIAG*5, x-RDIAG*6, x-RDIAG*7))
	return endPos
}

func (p *Piece) rookMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	endPos = appendIf(endPos, validPosList(x+RANK, x+RANK*2, x+RANK*3, x+RANK*4, x+RANK*5, x+RANK*6, x+RANK*7))
	endPos = appendIf(endPos, validPosList(x-RANK, x-RANK*2, x-RANK*3, x-RANK*4, x-RANK*5, x-RANK*6, x-RANK*7))
	endPos = appendIf(endPos, validPosList(x+FILE, x+FILE*2, x+FILE*3, x+FILE*4, x+FILE*5, x+FILE*6, x+FILE*7))
	endPos = appendIf(endPos, validPosList(x-FILE, x-FILE*2, x-FILE*3, x-FILE*4, x-FILE*5, x-FILE*6, x-FILE*7))
	return endPos
}

func (p *Piece) queenMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	endPos = append(endPos, p.bishopMove(pos)...)
	endPos = append(endPos, p.rookMove(pos)...)
	return endPos
}

func (p *Piece) kingMove(pos int8) [][]int8 {
	endPos := [][]int8{}
	x := int(pos)
	endPos = appendIf(endPos, validPosList(x+RANK))
	endPos = appendIf(endPos, validPosList(x-RANK))
	// if king is on proper square, return castling option
	if (p.Color == WHITE && pos == E1) || (p.Color == BLACK && pos == E8) {
		endPos = appendIf(endPos, validPosList(x-FILE, x-FILE*2))
		endPos = appendIf(endPos, validPosList(x+FILE, x+FILE*2))
	} else {
		endPos = appendIf(endPos, validPosList(x-FILE))
		endPos = appendIf(endPos, validPosList(x+FILE))
	}
	return endPos
}

// generate a list of possible squares a piece can move
// to given a starting position(square)
// we don't care about legal moves, we just want to generate
// the possible squares a piece can go, capture or no capture.
// The return value is a vector of vectors:
// Each sub-vector is a ray of moves, originating at the origin
// 'pos' and moving in a single direction (forward, back, left, right,
// left-diagonal and right-diagonal).
func (p *Piece) Move(pos int8) [][]int8 {
	switch p.Type {
	case PAWN:
		return p.pawnMove(pos)
	case KNIGHT:
		return p.knightMove(pos)
	case BISHOP:
		return p.bishopMove(pos)
	case ROOK:
		return p.rookMove(pos)
	case QUEEN:
		// this could be a slice join of rook and bishop moves
		return p.queenMove(pos)
	case KING:
		return p.kingMove(pos)
	default:
		fmt.Println("unknown piece type:", p.Type)
		break
	}

	return [][]int8{}
}

type Board struct {
	MoveCount int // track full moves, increment after every black move
	FiftyMove int // counter to track 50 move rule increment non-pawn/non-capture moves, otherwise reset

	ToMove    int8
	EnPassant int8

	KCastle [2]bool
	QCastle [2]bool

	Pieces  [COLOR_MAX][PIECES_MAX][]int8
	Squares [MAX_SQUARES]Piece
}

func (b *Board) FEN() *string {
	var fen string
	var empty int

	// FIELD : Piece Positions
	for _, rank := range []int8{RANK_8, RANK_7, RANK_6, RANK_5, RANK_4, RANK_3, RANK_2, RANK_1} {
		empty = 0
		for sq := rank; sq < rank+8; sq++ {
			p := b.Squares[sq]
			if p.Equal(&EmptyPiece) {
				empty++
			} else if empty != 0 {
				fen += strconv.Itoa(empty)
				empty = 0
			}
			fen += PieceSymbol[p.Color][p.Type]
		}
		if empty != 0 {
			fen += strconv.Itoa(empty)
		}
		fen += "/"
	}
	// we add a trailing slash, so remove it
	fen = strings.TrimSuffix(fen, "/")

	// FIELD : Side to Move
	if b.ToMove == WHITE {
		fen += " w "
	} else {
		fen += " b "
	}

	// FIELD : Castle Rights
	castle := ""
	if b.KCastle[WHITE] {
		castle += "K"
	}
	if b.QCastle[WHITE] {
		castle += "Q"
	}
	if b.KCastle[BLACK] {
		castle += "k"
	}
	if b.QCastle[BLACK] {
		castle += "q"
	}
	if castle == "" {
		fen += "-"
	} else {
		fen += castle
	}
	fen += " "

	// FIELD : En passant capture square
	fen += AlgebraicTable[b.EnPassant]
	fen += " "

	// FIELD : Half-Move (Fifty Move Rule) Clock
	fen += strconv.Itoa(b.FiftyMove)
	fen += " "

	// FIELD : Full move counter
	fen += strconv.Itoa(b.MoveCount)
	fen += " "

	return &fen
}

func (b *Board) GenMoves() []*Board {
	blist := []*Board{}

	for pos, piece := range b.Squares {
		if piece != EmptyPiece && piece.Color == b.ToMove {
			moveList := piece.Move(int8(pos))
			fmt.Printf("%s : %2d => %v : %d\n", PieceSymbol[piece.Color][piece.Type], pos, moveList, len(moveList))
		}
	}

	return blist
}


func (b *Board) IsValidMove(from int8, to int8) bool {
	mover := b.Squares[from]
	if mover == EmptyPiece {
		return false
	}

	dest := b.Squares[to]
	if dest != EmptyPiece && dest.Color == mover.Color {
		return false
	}
	
	return true
}

func iabs8(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

func (b *Board) Move(from int8, to int8) (*Board, error) {
	nb := *b

	if nb.IsValidMove(from, to) {
		piece1 := nb.Squares[from]
		piece2 := nb.Squares[to]
		if piece1.Color == BLACK {
			nb.MoveCount++
		}
		if piece1.Type == PAWN || piece2 != EmptyPiece {
			nb.FiftyMove++
		} else {
			nb.FiftyMove = 0
		}
		if piece1.Type == PAWN && iabs8(to - from) == RANK*2 {
			nb.EnPassant = from + ((to - from) / 2)
		}
		if piece1.Type == KING {
			nb.KCastle[piece1.Color] = false
			nb.QCastle[piece1.Color] = false
		}
		if piece1.Type == ROOK {
			if piece1.Color == WHITE {
			if from == A1 {
				nb.QCastle[WHITE] = false
			} else if from == H1 {
				nb.KCastle[WHITE] = false
			}
			} else {
			if from == A1 {
				nb.QCastle[BLACK] = false
			} else if from == H1 {
				nb.KCastle[BLACK] = false
			}
			}
		}
		//nb.Pieces[piece2.Color][piece2.Type]
		//nb.Pieces[piece1.Color][piece1.Type]
		nb.Squares[from] = EmptyPiece
		nb.Squares[to] = piece1
	} else {
		return b, errors.New(fmt.Sprintf("invalid move - %s => %s", AlgebraicTable[from], AlgebraicTable[to]))
	}

	return &nb, nil
}

func (b *Board) Place(p Piece, sq int8) *Board {
	nb := *b
	nb.Pieces[p.Color][p.Type] = append(nb.Pieces[p.Color][p.Type], sq)
	nb.Squares[sq] = p

	return &nb
}

func (b *Board) Debug() {
	for _, rank := range []int8{RANK_1, RANK_2, RANK_3, RANK_4, RANK_5, RANK_6, RANK_7, RANK_8} {
		fmt.Println("rank =", rank)
	}
}

func NewBoard() *Board {
	b := new(Board)
	b.ToMove = WHITE
	b.MoveCount = 1
	b.FiftyMove = 0
	b.EnPassant = NOSQ
	b.KCastle = [...]bool{true, true}
	b.QCastle = [...]bool{true, true}

	b.Pieces[WHITE][PAWN] = []int8{A2, B2, C2, D2, E2, F2, G2, H2}
	b.Pieces[BLACK][PAWN] = []int8{A7, B7, C7, D7, E7, F7, G7, H7}
	b.Pieces[WHITE][KNIGHT] = []int8{B1, G1}
	b.Pieces[BLACK][KNIGHT] = []int8{B8, G8}
	b.Pieces[WHITE][BISHOP] = []int8{C1, F1}
	b.Pieces[BLACK][BISHOP] = []int8{C8, F8}
	b.Pieces[WHITE][ROOK] = []int8{A1, H1}
	b.Pieces[BLACK][ROOK] = []int8{A8, H8}
	b.Pieces[WHITE][QUEEN] = []int8{D1}
	b.Pieces[BLACK][QUEEN] = []int8{D8}
	b.Pieces[WHITE][KING] = []int8{E1}
	b.Pieces[BLACK][KING] = []int8{E8}

	for c, _ := range b.Pieces {
		for p, _ := range b.Pieces[c] {
			for _, x := range b.Pieces[c][p] {
				b.Squares[x] = Piece{int8(c), int8(p)}
			}
		}
	}

	return b
}

func EmptyBoard() *Board {
	b := new(Board)
	b.ToMove = WHITE
	b.MoveCount = 1
	b.FiftyMove = 0
	b.EnPassant = NOSQ
	b.KCastle = [...]bool{true, true}
	b.QCastle = [...]bool{true, true}

	b.Pieces[WHITE][PAWN] = []int8{}
	b.Pieces[BLACK][PAWN] = []int8{}
	b.Pieces[WHITE][KNIGHT] = []int8{}
	b.Pieces[BLACK][KNIGHT] = []int8{}
	b.Pieces[WHITE][BISHOP] = []int8{}
	b.Pieces[BLACK][BISHOP] = []int8{}
	b.Pieces[WHITE][ROOK] = []int8{}
	b.Pieces[BLACK][ROOK] = []int8{}
	b.Pieces[WHITE][QUEEN] = []int8{}
	b.Pieces[BLACK][QUEEN] = []int8{}
	b.Pieces[WHITE][KING] = []int8{}
	b.Pieces[BLACK][KING] = []int8{}

	return b
}
