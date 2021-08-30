package chess

type iterateFn func(i int) error

// Iterates over the exclusive range (min{src,dest} , max{src,dest})
func iterBetween(src int, dest int, f iterateFn) (err error) {
	lo := min(src, dest) + 1
	hi := max(src, dest)

	for i := lo; i < hi; i++ {
		err = f(i)
	}
	return err
}

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
