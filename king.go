package main

type King struct {
	PieceInfo
}

func NewKing(p PieceInfo) *King {
	k := King{}
	k.PieceInfo = p
	return &k
}

func (k King) String() string {
	return "k"
}
