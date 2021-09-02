package main

type Bishop struct {
	PieceInfo
}

func (p *Bishop) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *Bishop) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Bishop) String() string {
	return diffPlayerPiece(bishop, p.player)
}

func (p *Bishop) threats(b Board) (threats []IPosn) {
	return iterThreats(p, b, diagonalDirs())
}

func (p *Bishop) validMoves(board Board) (dests []IPosn) {
	dests = iterMoves(p, board, diagonalDirs())
	return filterValidMoves(dests, p, board)
}
