package main

type Knight struct {
	PieceInfo
}

func NewKnight(p PieceInfo) *Knight {
	n := Knight{}
	n.PieceInfo = p
	return &n
}

func (n *Knight) moveIsValid(p Posn) bool {
	return true
}

func (n Knight) String() string {
	return n.formatStr("n")
}
