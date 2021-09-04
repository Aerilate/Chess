package main

var orthogonalDirs = []IPosn{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var diagonalDirs = []IPosn{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

func iterThreats(srcPiece Piece, board Board, incs []IPosn) (threats []IPosn) {
	for _, inc := range incs {
		curr := srcPiece.pieceInfo().IPosn
		curr = curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			threats = append(threats, curr)
			if !board.squareIsEmpty(curr) { // piece at edge of threat
				break
			}
			curr = curr.add(inc)
		}
	}
	return threats
}

func iterMoves(srcPiece Piece, board Board, incs []IPosn) (dests []IPosn) {
	for _, inc := range incs {
		curr := srcPiece.pieceInfo().IPosn
		curr = curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			if !board.squareIsEmpty(curr) {
				if areEnemies(board.at(curr), srcPiece) { // opponent piece can be captured
					dests = append(dests, curr)
				}
				break // line of sight ends
			}
			dests = append(dests, curr)
			curr = curr.add(inc)
		}
	}
	return dests
}
