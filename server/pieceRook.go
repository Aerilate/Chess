package main

type Rook struct {
	PieceInfo
}

func (p *Rook) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *Rook) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Rook) String() string {
	return diffPlayerPiece(rook, p.player)
}

func (p *Rook) threats(board Board) (threats []IPosn) {
	return iterThreats(p, board, orthogonalDirs())
}

func (p *Rook) validMoves(board Board) (dests []IPosn) {
	dests = iterMoves(p, board, orthogonalDirs())
	return filterValidMoves(dests, p, board)
}
