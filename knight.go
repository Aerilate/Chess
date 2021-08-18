package main

type Knight struct {
	PieceInfo
}

func NewKnight(p PieceInfo) *Knight {
	n := Knight{}
	n.PieceInfo = p
	return &n
}

func (k Knight) String() string {
	return "n"
}
