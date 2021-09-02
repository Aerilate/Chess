package main

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
	threats(b Board) (threats []IPosn)
	validMoves(b Board) (moves []IPosn)

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

func filterValidMoves(dests []IPosn, piece Piece, b Board) []IPosn {
	src := piece.pieceInfo().IPosn
	player := piece.pieceInfo().player

	destBlocked := func(p IPosn) bool {
		return *b.at(p) != nil && (*b.at(p)).pieceInfo().player == player
	}

	kingChecked := func(p IPosn) bool {
		copy := b.shallowCopy()
		*copy.at(p) = *copy.at(src)
		(*copy.at(p)).updatePosn(p)
		*copy.at(src) = nil
		return copy.kingUnderCheck(player)
	}

	return filter(dests, func(p IPosn) bool {
		return moveInBounds(p) && !destBlocked(p) && !kingChecked(p)
	})
}
