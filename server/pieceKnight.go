package main

type Knight struct {
	PieceInfo
}

func (p *Knight) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *Knight) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Knight) String() string {
	return diffPlayerPiece(knight, p.player)
}

func (p *Knight) checkMove(b Board, threats ChecksBoard, dest IPosn) error {
	if (abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 2) ||
		(abs(dest.i-p.i) == 2 && abs(dest.j-p.j) == 1) {
		return nil
	}
	return InvalidMove{"Knight can't move there."}
}

func (p *Knight) threats(b Board) (threats []IPosn) {
	threats = append(threats, IPosn{p.i + 2, p.j + 1}, IPosn{p.i + 2, p.j - 1})
	threats = append(threats, IPosn{p.i - 2, p.j + 1}, IPosn{p.i - 2, p.j - 1})
	threats = append(threats, IPosn{p.i + 1, p.j + 2}, IPosn{p.i + 1, p.j - 2})
	threats = append(threats, IPosn{p.i - 1, p.j + 2}, IPosn{p.i - 1, p.j - 2})
	return threats
}
