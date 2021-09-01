package main

type Move struct {
	src  IPosn
	dest IPosn
}

type Game interface {
	move(src IPosn, dest IPosn) error
	calcValidMoves() []Move
	getActivePlayer() int
	gameIsOver() bool
	recentMove() Move
	String() string
}

func NewGame() Game {
	return NewGameState()
}
