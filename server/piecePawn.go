package main

type Pawn struct {
	PieceInfo
}

func (p Pawn) pieceInfo() PieceInfo {
	return p.PieceInfo
}

func (p Pawn) updatePosn(posn IPosn) {
	p.PieceInfo.IPosn = posn
}

func (p Pawn) String() string {
	return diffPlayerPiece(pawn, p.player)
}

func (p Pawn) threats(board Board) (threats []IPosn) {
	moveDir := moveDirection(p.player)
	return []IPosn{{p.i + moveDir, p.j - 1}, {p.i + moveDir, p.j + 1}}
}

func pawnHomeRow(player int) int {
	if player == Player1 {
		return 6
	}
	return 1
}

// maps 1->-1 and 2->1
func moveDirection(player int) int {
	return player*2 - 3
}

func (p Pawn) validDests(board Board) (dests []IPosn) {
	pawnRow := pawnHomeRow(p.player)
	moveDir := moveDirection(p.player)

	oneAhead := IPosn{p.i + moveDir, p.j}
	if p.i == pawnRow && board.squareIsEmpty(oneAhead) {
		dests = append(dests, IPosn{p.i + 2*moveDir, p.j}) // can go two ahead
	}
	dests = append(dests, oneAhead)

	// capturable squares
	diagLeft := IPosn{p.i + moveDir, p.j - 1}
	diagRight := IPosn{p.i + moveDir, p.j + 1}
	if moveInBounds(diagLeft) && areEnemies(p, board.at(diagLeft)) {
		dests = append(dests, diagLeft)
	}
	if moveInBounds(diagRight) && areEnemies(p, board.at(diagRight)) {
		dests = append(dests, diagRight)
	}
	return filterValidMoves(dests, p, board)
}
