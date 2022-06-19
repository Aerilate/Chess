package main

type Game interface {
	ValidMoves() map[string][]string
	Move(move Move)
	ActivePlayer() Player
	Checked() Player
	IsOver() bool
	Fen() string
	String() string
}

type Move struct {
	src  StdPosn
	dest StdPosn
}

func NewGame() Game {
	return NewGameState()
}
