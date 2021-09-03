package main

type Move struct {
	src  StdPosn
	dest StdPosn
}

type Game interface {
	validMoves() map[string][]string
	move(move Move)
	ActivePlayer() int
	isOver() bool
	fen() string
}

func NewGame() Game {
	return NewGameState()
}
