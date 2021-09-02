package main

type Move struct {
	src  StdPosn
	dest StdPosn
}

type Game interface {
	validMoves() map[string][]string
	move(move Move)
	lastMove() Move
	getActivePlayer() int
	isOver() bool
	String() string
}

func NewGame() Game {
	return NewGameState()
}
