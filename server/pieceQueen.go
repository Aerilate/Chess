package main

type Queen struct {
	PieceInfo
}

func (p Queen) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p Queen) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Queen) String() string {
	return diffPlayerPiece(queen, p.player)
}

func (p Queen) threats(board Board) (threats []IPosn) {
	orthogonalThreats := iterThreats(p, board, orthogonalDirs)
	diagonalThreats := iterThreats(p, board, diagonalDirs)
	return append(orthogonalThreats, diagonalThreats...)
}

func (p Queen) validMoves(board Board) (dests []IPosn) {
	orthogonalDests := iterMoves(p, board, orthogonalDirs)
	diagonalDests := iterMoves(p, board, diagonalDirs)
	dests = append(orthogonalDests, diagonalDests...)
	return filterValidMoves(dests, p, board)
}
