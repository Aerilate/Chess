package main

type InvalidMove struct {
	s string
}

func (e InvalidMove) Error() string {
	return "Invalid move, " + e.s + " Try again: "
}

type NotYourPiece struct{}

func (e NotYourPiece) Error() string {
	return "That isn't your piece. Try again: "
}
