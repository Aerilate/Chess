package main

type King struct {
	PieceInfo
}

func (p *King) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *King) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p King) String() string {
	return diffPlayerPiece(king, p.player)
}

func (p *King) checkMove(b Board, threats ChecksBoard, dest IPosn) error {
	if abs(dest.i-p.i) <= 1 && abs(dest.j-p.j) <= 1 && threats.squareIsSafe(dest) {
		return nil
	}
	return InvalidMove{"King can't move there."}
}

func (p *King) threats(b Board) (threats []IPosn) {
	for i := p.i - 1; i <= p.i+1; i++ {
		for j := p.j - 1; j <= p.j+1; j++ {
			threats = append(threats, IPosn{i, j})
		}
	}
	return threats
}
