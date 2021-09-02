package main

type Move struct {
	src  StdPosn
	dest StdPosn
}

type Game interface {
	validMoves() map[StdPosn][]StdPosn
	move(move Move)
	lastMove() Move
	getActivePlayer() int
	gameIsOver() bool
	String() string
}

func NewGame() Game {
	return NewGameState()
}
