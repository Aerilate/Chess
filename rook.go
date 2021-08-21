package main

type Rook struct {
	PieceInfo
}

func NewRook(p PieceInfo) *Rook {
	r := Rook{}
	r.PieceInfo = p
	return &r
}

func (r *Rook) moveIsValid(p Posn) bool {
	return true
}

func (r Rook) String() string {
	return r.formatStr("r")
}
