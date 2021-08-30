package chess

import "strings"

const (
	pawn   = "p"
	knight = "n"
	bishop = "b"
	rook   = "r"
	queen  = "q"
	king   = "k"
)

type Piece interface {
	checkMove(board Board, threats ChecksBoard, dest IPosn) (err error)
	threats(b Board) (threats []IPosn)

	pieceInfo() PieceInfo
	updatePosn(posn IPosn)
	String() string
}

type PieceInfo struct {
	player int
	IPosn
}

func areEnemies(p1 Piece, p2 Piece) bool {
	return p2 != nil && p1.pieceInfo().player != p2.pieceInfo().player
}

// fn to differentiate player pieces
func diffPlayerPiece(s string, player int) string {
	if player == Player1 {
		return strings.ToUpper(s)
	}
	return s
}
