package main

type Pawn struct {
	PieceInfo
}

func (p *Pawn) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p *Pawn) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Pawn) String() string {
	return diffPlayerPiece(pawn, p.player)
}

func (p *Pawn) checkMove(board Board, threats ChecksBoard, dest IPosn) error {
	pawnRow := pawnHomeRow(p.player)
	moveDir := moveDirection(p.player)

	if dest.j == p.j { // move straight ahead
		if board[p.i+moveDir][p.j] != nil {
			return InvalidMove{"Piece in the way!"}
		}

		if p.i == pawnRow && dest.i == p.i+moveDir*2 { // move two squares from home row
			if board[p.i+moveDir*2][p.j] == nil {
				return nil
			}
			return InvalidMove{"Piece in the way!"}
		} else if dest.i == p.i+moveDir {
			return nil
		}
	} else if dest.j == p.j-1 || dest.j == p.j+1 { // diagonal move
		if dest.i == p.i+moveDir && areEnemies(p, *board.at(dest)) {
			return nil
		}
	}
	return InvalidMove{"Pawn can't move there."}
}

func (p *Pawn) threats(b Board) (threats []IPosn) {
	moveDir := moveDirection(p.player)
	return []IPosn{{p.i + moveDir, p.j - 1}, {p.i + moveDir, p.j + 1}}
}
