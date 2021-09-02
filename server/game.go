package main

type Move struct {
	src  StdPosn
	dest StdPosn
}

type Game interface {
	validMoves() map[string][]string
	move(move Move)
	fen() string
	getActivePlayer() int
	isOver() bool
}

func NewGame() Game {
	return NewGameState()
}
