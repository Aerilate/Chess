package main

type Gameable interface {
	ActivePlayer() Player
	Checked() bool
	IsOver() bool

	Move(string) error
	Fen() string
	String() string
}

func NewGame() Game {
	return NewGameState()
}