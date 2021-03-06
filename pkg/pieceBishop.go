package main

type Bishop struct {
	PieceInfo
}

func (p Bishop) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p Bishop) copy(info PieceInfo) Piece {
	return Bishop{info}
}

func (p Bishop) String() string {
	return diffPlayerPiece(bishop, p.player)
}

func (p Bishop) threats(board Board) (threats []IPosn) {
	return iterThreats(p, board, diagonalDirs)
}

func (p Bishop) validDests(board Board) (dests []IPosn) {
	dests = iterMoves(p, board, diagonalDirs)
	return filterValidMoves(dests, p, board)
}
