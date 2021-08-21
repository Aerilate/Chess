package main

type Bishop struct {
	PieceInfo
}

func NewBishop(p PieceInfo) *Bishop {
	b := Bishop{}
	b.PieceInfo = p
	return &b
}

func (b *Bishop) moveIsValid(p Posn) bool {
	return true
}

func (b Bishop) String() string {
	return b.formatStr("b")
}
