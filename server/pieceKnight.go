package main

type Knight struct {
	PieceInfo
}

func (p Knight) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p Knight) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Knight) String() string {
	return diffPlayerPiece(knight, p.player)
}

func (p Knight) threats(board Board) (threats []IPosn) {
	threats = append(threats, IPosn{p.i + 2, p.j + 1}, IPosn{p.i + 2, p.j - 1})
	threats = append(threats, IPosn{p.i - 2, p.j + 1}, IPosn{p.i - 2, p.j - 1})
	threats = append(threats, IPosn{p.i + 1, p.j + 2}, IPosn{p.i + 1, p.j - 2})
	threats = append(threats, IPosn{p.i - 1, p.j + 2}, IPosn{p.i - 1, p.j - 2})
	return threats
}

func (p Knight) validMoves(board Board) (dests []IPosn) {
	dests = p.threats(board)
	return filterValidMoves(dests, p, board)
}
