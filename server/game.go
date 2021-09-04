package main

type Game interface {
	ValidMoves() map[string][]string
	Move(move Move)
	ActivePlayer() int
	IsOver() bool
	Fen() string
}

type Move struct {
	src  StdPosn
	dest StdPosn
}

func NewGame() Game {
	return NewGameState()
}
