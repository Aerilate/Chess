package main

import (
	"strings"
)

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
	return p1 != nil && p2 != nil && p1.pieceInfo().player != p2.pieceInfo().player
}

// fn to differentiate player pieces
func diffPlayerPiece(s string, player int) string {
	if player == Player1 {
		return strings.ToUpper(s)
	}
	return s
}

func filterValidMoves(dests []IPosn, piece Piece, board Board) []IPosn {
	src := piece.pieceInfo().IPosn
	player := piece.pieceInfo().player

	destCapturable := func(dest IPosn) bool {
		return board.squareIsEmpty(dest) || areEnemies(*board.at(dest), piece)
	}

	kingChecked := func(dest IPosn) bool {
		copy := board.shallowCopy()
		*copy.at(dest) = *copy.at(src)
		// do not update piece posn since we're working with a shallow copy of the piece
		*copy.at(src) = nil
		return copy.kingUnderCheck(player)
	}

	return filter(dests, func(p IPosn) bool {
		return moveInBounds(p) && destCapturable(p) && !kingChecked(p)
	})
}
