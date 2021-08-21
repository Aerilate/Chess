package main

type Queen struct {
	PieceInfo
}

func NewQueen(p PieceInfo) *Queen {
	q := Queen{}
	q.PieceInfo = p
	return &q
}

func (q *Queen) moveIsValid(p Posn) bool {
	return true
}

func (q Queen) String() string {
	return q.formatStr("q")
}
