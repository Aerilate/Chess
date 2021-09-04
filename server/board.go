package main

const BoardSize = 8

type Board interface {
	at(p IPosn) Piece
	squareIsEmpty(p IPosn) bool
	setSquare(dest IPosn, piece Piece)
	pieces() []Piece
	deepCopy() Board
	fen() string
}

func NewBoard() Board {
	return NewBoardImp()
}

func moveInBounds(p IPosn) bool {
	return 0 <= p.i && p.i < BoardSize && 0 <= p.j && p.j < BoardSize
}
