package main

type Queen struct {
	PieceInfo
}

func (p *Queen) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *Queen) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Queen) String() string {
	return diffPlayerPiece(queen, p.player)
}

func (p *Queen) checkMove(board Board, threats ChecksBoard, dest IPosn) (err error) {
	ortho := p.orthogonalThreats(board, dest)
	diag := p.diagonalThreats(board, dest)
	if ortho == nil || diag == nil {
		return nil
	}
	return InvalidMove{"Queen can't move there."}
}

func (p *Queen) threats(b Board) (threats []IPosn) {
	orthogonalThreats := iterThreats(p, b, orthogonalDirs())
	diagonalThreats := iterThreats(p, b, diagonalDirs())
	return append(orthogonalThreats, diagonalThreats...)
}

func (p *Queen) validMoves(board Board, threats ChecksBoard) (dests []IPosn) {
	return nil
}
