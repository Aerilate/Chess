package main

type Bishop struct {
	PieceInfo
}

func NewBishop(p PieceInfo) *Bishop {
	b := Bishop{}
	b.PieceInfo = p
	return &b
}

func (b Bishop) String() string {
	return "b"
}
