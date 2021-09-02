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

func (p *Rook) checkMove(board Board, threats ChecksBoard, dest IPosn) (err error) {
	err = p.orthogonalThreats(board, dest)
	if err == nil {
		return nil
	}
	return InvalidMove{"Rook can't move there."}
}

func (p *PieceInfo) orthogonalThreats(board Board, dest IPosn) (err error) {
	if abs(dest.j-p.j) == 0 {
		f := func(i int) (err error) {
			if board[i][p.j] != nil {
				return InvalidMove{"Piece in the way."}
			}
			return nil
		}
		return iterBetween(p.i, dest.i, f)
	} else if abs(dest.i-p.i) == 0 {
		f := func(i int) (err error) {
			if board[p.i][i] != nil {
				return InvalidMove{"Piece in the way."}
			}
			return nil
		}
		return iterBetween(p.j, dest.j, f)
	}
	return InvalidMove{"Not an orthogonal move!"}
}

func (p *Rook) threats(b Board) (threats []IPosn) {
	return iterThreats(p, b, orthogonalDirs())
}

func (p *Rook) validMoves(board Board, threats ChecksBoard) (dests []IPosn) {
	return nil
}
