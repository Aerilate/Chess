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

func (p *Piece) canCapture(other *Piece) bool {
	if other == nil {
		return false
	}
	return p.player != other.player && other.pieceType != king
}

func (p *Piece) String() string {
	if p.player == Player1 {
		return strings.ToUpper(p.pieceType)
	}
	return p.pieceType
}

func (p *Piece) checkMove(vc *ViewController, dest Posn) (err error) {
	switch p.pieceType {
	case pawn:
		err = p.checkPawnMove(&vc.Board, dest)
	case knight:
		err = p.checkKnightMove(&vc.Board, dest)
	case bishop:
	case rook:
	case queen:
	case king:
	}
	return
}

func (p *Piece) checkPawnMove(board *Board, dest Posn) error {
	if (p.player == Player1 && p.i == 6 && dest.i == p.i-2) ||
		(p.player == Player2 && p.i == 1 && dest.i == p.i+2) { // move two squares from home row
		return nil
	} else if (p.player == Player1 && dest.i == p.i-1) ||
		(p.player == Player2 && dest.i == p.i+1) {
		return nil
	}

	if dest.j != p.j-1 && dest.j != p.j && dest.j != p.j+1 {
		return InvalidMove{"Pawn can't move more than one square horizontally."}
	}
	return InvalidMove{"Pawn can't move there vertically."}
}

func (p *Piece) checkKnightMove(board *Board, dest Posn) error {
	if (abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 2) ||
		(abs(dest.i-p.i) == 2 && abs(dest.j-p.j) == 1) {
		return nil
	}
	return InvalidMove{"Knight can't move there."}
}

func (p *Piece) checkRookMove(board *Board, dest Posn) error {
	if abs(dest.j-p.j) == 0 {
		for i := p.i + 1; i < dest.i; i += mag(dest.i - p.i) {
			if board.board[i][p.j] != nil {
				return InvalidMove{"Piece in the way."}
			}
		}
		if board.at(dest) == nil || p.canCapture(*board.at(dest)) {
			return nil
		}
		return InvalidMove{"Piece in the way."}
	} else if abs(dest.i-p.i) == 0 {
		for j := p.j + 1; j < dest.j; j += mag(dest.j - p.j) {
			if board.board[p.i][j] != nil {
				return InvalidMove{"Piece in the way."}
			}
		}
		if board.at(dest) == nil || p.canCapture(*board.at(dest)) {
			return nil
		}
		return InvalidMove{"Piece in the way."}
	}
	return InvalidMove{"Rook can't move there."}
}

func (p *Piece) checkKingMove(board *Board, dest Posn) error {
	if abs(dest.i-p.i) == 1 && abs(dest.j-p.j) == 1 {
		return nil
	}
	return InvalidMove{"King can't move there."}
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func mag(n int) int {
	if n == 0 {
		return 0
	}
	return n / abs(n)
}
