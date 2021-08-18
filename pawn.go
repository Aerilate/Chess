package main

type Pawn struct {
	posn
}

func (p Pawn) String() string {
	return "p"
}
