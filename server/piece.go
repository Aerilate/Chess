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

func NewPiece(pieceType string, info PieceInfo) (p Piece) {
	switch pieceType {
	case pawn:
		p = &Pawn{info}
	case knight:
		p = &Knight{info}
	case bishop:
		p = &Bishop{info}
	case rook:
		p = &Rook{info}
	case queen:
		p = &Queen{info}
	case king:
		p = &King{info}
	}
	return p
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
		return kingUnderCheck(copy, player)
	}

	return filter(dests, func(p IPosn) bool {
		return moveInBounds(p) && destCapturable(p) && !kingChecked(p)
	})
}
