package main

type Move struct {
	src  StdPosn
	dest StdPosn
}

type Game interface {
	move(move Move) error
	validMoves() map[StdPosn][]StdPosn
	getActivePlayer() int
	gameIsOver() bool
	lastMove() Move
	String() string
}

func NewGame() Game {
	return NewGameState()
}
