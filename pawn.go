package main

type Pawn struct {
	PieceInfo
}

func NewPawn(p PieceInfo) *Pawn {
	pw := Pawn{}
	pw.PieceInfo = p
	return &pw
}

func (p Pawn) String() string {
	return p.formatStr("p")
}
