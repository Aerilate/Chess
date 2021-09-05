package main

import "strings"

var orthogonalDirs = []IPosn{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var diagonalDirs = []IPosn{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

func iterThreats(srcPiece Piece, board Board, incs []IPosn) (threats []IPosn) {
	for _, inc := range incs {
		curr := srcPiece.pieceInfo().IPosn
		curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			threats = append(threats, curr)
			if !board.squareIsEmpty(curr) { // piece at end of line of sight
				break
			}
			curr.add(inc)
		}
	}
	return threats
}

func iterMoves(srcPiece Piece, board Board, incs []IPosn) (dests []IPosn) {
	for _, inc := range incs {
		dest := srcPiece.pieceInfo().IPosn
		dest.add(inc) // exclude piece position itself
		for moveInBounds(dest) {
			if areFriends(board.at(dest), srcPiece) {
				break // line of sight excludes dest
			}

			dests = append(dests, dest)
			if areEnemies(board.at(dest), srcPiece) {
				break // line of sight includes dest but ends
			}
			dest.add(inc)
		}
	}
	return dests
}

func filterValidMoves(dests []IPosn, piece Piece, board Board) (result []IPosn) {
	src := piece.pieceInfo().IPosn
	player := piece.pieceInfo().player

	destCapturable := func(dest IPosn) bool {
		return board.squareIsEmpty(dest) || areEnemies(board.at(dest), piece)
	}

	kingChecked := func(dest IPosn) bool {
		boardCopy := board.deepCopy()
		boardCopy.setSquare(dest, boardCopy.at(src))
		boardCopy.setSquare(src, nil)
		return kingUnderCheck(boardCopy, player)
	}

	// apply filter
	for _, dest := range dests {
		if moveInBounds(dest) && destCapturable(dest) && !kingChecked(dest) {
			result = append(result, dest)
		}
	}
	return result
}

func areFriends(p1 Piece, p2 Piece) bool {
	return p1 != nil && p2 != nil && p1.pieceInfo().player == p2.pieceInfo().player
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
