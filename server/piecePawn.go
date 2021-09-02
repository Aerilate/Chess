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

func (p *Pawn) threats(board Board) (threats []IPosn) {
	moveDir := moveDirection(p.player)
	return []IPosn{{p.i + moveDir, p.j - 1}, {p.i + moveDir, p.j + 1}}
}

func (p *Pawn) validMoves(board Board) (dests []IPosn) {
	pawnRow := pawnHomeRow(p.player)
	moveDir := moveDirection(p.player)

	oneAhead := IPosn{p.i + moveDir, p.j}
	if p.i == pawnRow {
		if board.at(oneAhead) == nil { // can go two ahead
			dests = append(dests, oneAhead.add(IPosn{1, 0}))
		}
	}
	dests = append(dests, oneAhead)

	// capturable squares
	diagLeft := IPosn{p.i + moveDir, p.j - 1}
	diagRight := IPosn{p.i + moveDir, p.j + 1}
	if moveInBounds(diagLeft) && areEnemies(*board.at(diagLeft), p) {
		dests = append(dests, diagLeft)
	}
	if moveInBounds(diagRight) && areEnemies(*board.at(diagRight), p) {
		dests = append(dests, diagRight)
	}
	return filterValidMoves(dests, p, board)
}
