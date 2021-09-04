package main

type King struct {
	PieceInfo
}

func (p King) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p King) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p King) String() string {
	return diffPlayerPiece(king, p.player)
}

func (p King) threats(board Board) (threats []IPosn) {
	for i := p.i - 1; i <= p.i+1; i++ {
		for j := p.j - 1; j <= p.j+1; j++ {
			threats = append(threats, IPosn{i, j})
		}
	}
	return threats
}

func (p King) validDests(board Board) (dests []IPosn) {
	for i := p.i - 1; i <= p.i+1; i++ {
		for j := p.j - 1; j <= p.j+1; j++ {
			dests = append(dests, IPosn{i, j})
		}
	}
	return filterValidMoves(dests, p, board)
}
