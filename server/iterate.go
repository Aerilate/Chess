package main

func orthogonalDirs() []IPosn {
	return []IPosn{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
}

func diagonalDirs() []IPosn {
	return []IPosn{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
}

func iterThreats(p Piece, b Board, incs []IPosn) (threats []IPosn) {
	for _, inc := range incs {
		curr := p.pieceInfo().IPosn
		curr = curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			threats = append(threats, curr)
			if *b.at(curr) != nil { // piece at edge of threat
				break
			}
			curr = curr.add(inc)
		}
	}
	return threats
}

func iterMoves(p Piece, b Board, incs []IPosn) (dests []IPosn) {
	for _, inc := range incs {
		curr := p.pieceInfo().IPosn
		curr = curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			if *b.at(curr) != nil {
				destPlayer := (*b.at(curr)).pieceInfo().player
				if destPlayer != p.pieceInfo().player { // opponent piece can be captured
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
