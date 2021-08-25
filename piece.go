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
	Posn
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

func (p *Piece) checkMove(vc *ViewController, dest Posn) (err error) {
	if *vc.at(dest) != nil && p.player == (*vc.at(dest)).player { // check if dest occupied by own piece
		return InvalidMove{"Coordinate " + dest.String() + " is occupied by your own piece!"}
	}

	switch p.pieceType {
	case pawn:
		err = p.checkPawnMove(vc.Board, dest)
	case knight:
		err = p.checkKnightMove(dest)
	case bishop:
		err = p.checkBishopMove(vc.Board, dest)
	case rook:
		err = p.checkRookMove(vc.Board, dest)
	case queen:
		err = p.checkQueenMove(vc.Board, dest)
	case king:
		err = p.checkKingMove(dest)
	}
	return err
}

func (p *Piece) checkPawnMove(board Board, dest Posn) error {
	pawnRow := pawnHomeRow(p.player)
	moveDir := moveDirection(p.player)

	if dest.j == p.j { // move straight ahead
		if board[dest.i+moveDir][dest.j] != nil {
			return InvalidMove{"Piece in the way!"}
		}

		if p.i == pawnRow && dest.i == p.i+moveDir*2 { // move two squares from home row
			if board[dest.i+moveDir*2][dest.j] == nil {
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

func (p *Piece) checkKnightMove(dest Posn) error {
	if (abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 2) ||
		(abs(dest.i-p.i) == 2 && abs(dest.j-p.j) == 1) {
		return nil
	}
	return InvalidMove{"Knight can't move there."}
}

func (p *Piece) checkBishopMove(board Board, dest Posn) error {
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

func (p *Piece) checkRookMove(board Board, dest Posn) (err error) {
	if abs(dest.j-p.j) == 0 {
		f := func(i int) (err error) {
			if board[i][p.j] != nil {
				return InvalidMove{"Piece in the way."}
			}
			return nil
		}
		return iterOverExclu(p.i, dest.i, f)
	} else if abs(dest.i-p.i) == 0 {
		f := func(i int) (err error) {
			if board[p.i][i] != nil {
				return InvalidMove{"Piece in the way."}
			}
			return nil
		}
		return iterOverExclu(p.j, dest.j, f)
	}
	return InvalidMove{"Rook can't move there."}
}

func (p *Piece) checkQueenMove(board Board, dest Posn) error {
	lateral := p.checkRookMove(board, dest)
	diag := p.checkBishopMove(board, dest)
	if lateral == nil || diag == nil {
		return nil
	}
	return InvalidMove{"Queen can't move there."}
}

func (p *Piece) checkKingMove(dest Posn) error {
	if abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 1 {
		return nil
	}
	return InvalidMove{"King can't move there."}
}

type iterateFn func(i int) error

func iterOverExclu(src int, dest int, f iterateFn) (err error) {
	lo := min(src, dest) + 1
	hi := max(src, dest)

	for i := lo; i < hi; i++ {
		err = f(i)
	}
	return err
}
