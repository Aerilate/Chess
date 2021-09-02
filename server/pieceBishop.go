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

func (p *Bishop) checkMove(board Board, threats ChecksBoard, dest IPosn) (err error) {
	err = p.diagonalThreats(board, dest)
	if err == nil {
		return nil
	}
	return InvalidMove{"Rook can't move there."}
}

func (p *PieceInfo) diagonalThreats(board Board, dest IPosn) (err error) {
	iDiff := dest.i - p.i
	jDiff := dest.j - p.j
	if abs(iDiff) == abs(jDiff) {
		for n := 1; n < abs(iDiff); n += 1 {
			if board[p.i+n*mag(iDiff)][p.j+n*mag(jDiff)] != nil {
				return InvalidMove{"Piece in the way."}
			}
		}
		return nil
	}
	return InvalidMove{"Not a diagonal move!"}
}

func (p *Bishop) threats(b Board) (threats []IPosn) {
	return iterThreats(p, b, diagonalDirs())
}

func (p *Bishop) validMoves(board Board, threats ChecksBoard) (dests []IPosn) {
	return nil
}
