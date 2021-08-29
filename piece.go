package main

import (
	"strings"
)

const (
	pawn   = "p"
	knight = "n"
	bishop = "b"
	rook   = "r"
	queen  = "q"
	king   = "k"
)

type Piece struct {
	pieceType string
	player    int
	IPosn
}

func (p *Piece) String() string {
	if p.player == Player1 {
		return strings.ToUpper(p.pieceType)
	}
	return p.pieceType
}

func (p *Piece) isEnemyTo(other *Piece) bool {
	return other != nil && p.player != other.player
}

func (p *Piece) checkMove(board Board, threats ChecksBoard, dest IPosn) (err error) {
	switch p.pieceType {
	case pawn:
		err = p.checkPawnMove(board, dest)
	case knight:
		err = p.checkKnightMove(dest)
	case bishop:
		err = p.checkBishopMove(board, dest)
	case rook:
		err = p.checkRookMove(board, dest)
	case queen:
		err = p.checkQueenMove(board, dest)
	case king:
		err = p.checkKingMove(threats, dest)
	}
	return err
}

func (p *Piece) checkPawnMove(board Board, dest IPosn) error {
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
		if dest.i == p.i+moveDir && p.isEnemyTo(*board.at(dest)) {
			return nil
		}
	}
	return InvalidMove{"Pawn can't move there."}
}

func (p *Piece) checkKnightMove(dest IPosn) error {
	if (abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 2) ||
		(abs(dest.i-p.i) == 2 && abs(dest.j-p.j) == 1) {
		return nil
	}
	return InvalidMove{"Knight can't move there."}
}

func (p *Piece) checkBishopMove(board Board, dest IPosn) error {
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
	return InvalidMove{"Bishop can't move there."}
}

func (p *Piece) checkRookMove(board Board, dest IPosn) (err error) {
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
	return InvalidMove{"Rook can't move there."}
}

func (p *Piece) checkQueenMove(board Board, dest IPosn) error {
	lateral := p.checkRookMove(board, dest)
	diag := p.checkBishopMove(board, dest)
	if lateral == nil || diag == nil {
		return nil
	}
	return InvalidMove{"Queen can't move there."}
}

func (p *Piece) checkKingMove(threats ChecksBoard, dest IPosn) error {
	if abs(dest.i-p.i) <= 1 && abs(dest.j-p.j) <= 1 && threats.squareIsSafe(dest) {
		return nil
	}
	return InvalidMove{"King can't move there."}
}

type iterateFn func(i int) error

func iterBetween(src int, dest int, f iterateFn) (err error) {
	lo := min(src, dest) + 1
	hi := max(src, dest)

	for i := lo; i < hi; i++ {
		err = f(i)
	}
	return err
}

func (p *Piece) threats(b Board) (threats []IPosn) {
	switch p.pieceType {
	case pawn:
		threats = p.pawnThreats()
	case knight:
		threats = p.knightThreats()
	case bishop:
		threats = p.bishopThreats(b)
	case rook:
		threats = p.rookThreats(b)
	case queen:
		threats = p.queenThreats(b)
	case king:
		threats = p.kingThreats()
	}
	return threats
}

func (p *Piece) pawnThreats() (threats []IPosn) {
	moveDir := moveDirection(p.player)
	return []IPosn{{p.i + moveDir, p.j - 1}, {p.i + moveDir, p.j + 1}}
}

func (p *Piece) knightThreats() (threats []IPosn) {
	threats = append(threats, IPosn{p.i + 2, p.j + 1}, IPosn{p.i + 2, p.j - 1})
	threats = append(threats, IPosn{p.i - 2, p.j + 1}, IPosn{p.i - 2, p.j - 1})
	threats = append(threats, IPosn{p.i + 1, p.j + 2}, IPosn{p.i + 1, p.j - 2})
	threats = append(threats, IPosn{p.i - 1, p.j + 2}, IPosn{p.i - 1, p.j - 2})
	return threats
}

func (p *Piece) bishopThreats(b Board) (threats []IPosn) {
	incs := []IPosn{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	return p.iterThreats(b, incs)
}

func (p *Piece) rookThreats(b Board) (threats []IPosn) {
	incs := []IPosn{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	return p.iterThreats(b, incs)
}

func (p *Piece) queenThreats(b Board) (threats []IPosn) {
	return append(p.bishopThreats(b), p.rookThreats(b)...)
}

func (p *Piece) kingThreats() (threats []IPosn) {
	for i := p.i - 1; i <= p.i+1; i++ {
		for j := p.j - 1; j <= p.j+1; j++ {
			threats = append(threats, IPosn{i, j})
		}
	}
	return threats
}

func (p *Piece) iterThreats(b Board, incs []IPosn) (threats []IPosn) {
	for _, inc := range incs {
		curr := p.IPosn
		curr = curr.add(inc) // exclude piece position itself
		for moveInBounds(curr) {
			threats = append(threats, curr)
			if *b.at(curr) != nil { // piece at edge of threat
				break
			}
			curr = curr.add(inc)
		}
	}
	return threats
}
