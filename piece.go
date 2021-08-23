package main

import "strings"

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

func (p *Piece) checkMove(vc *ViewController, dest Posn) (err error) {
	switch p.pieceType {
	case pawn:
		err = p.checkPawnMove(vc, dest)
	case knight:
	case bishop:
	case rook:
	case queen:
	case king:
	}
	return
}

func (p *Piece) checkPawnMove(vc *ViewController, dest Posn) error {
	if p.player != vc.activePlayer {
		return NotYourPiece{}
	}

	if dest.x != p.x-1 && dest.x != p.x && dest.x != p.x+1 {
		return InvalidMove{"Pawn can't move more than one unit horizontally."}
	}

	if (p.player == Player1 && dest.y == p.y-1) || (p.player == Player2 && dest.y == p.y+1) {
		return nil
	} else {
		return InvalidMove{"Pawn can't move there vertically."}
	}
}
