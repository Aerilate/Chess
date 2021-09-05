package main

import "strconv"

type BoardImp [BoardSize][BoardSize]Piece

func NewBoardImp() (b *BoardImp) {
	return &BoardImp{}
}

func (b BoardImp) at(p IPosn) Piece {
	return b[p.i][p.j]
}

func (b BoardImp) squareIsEmpty(p IPosn) bool {
	return b[p.i][p.j] == nil
}

func (b *BoardImp) setSquare(dest IPosn, piece Piece) {
	if piece == nil {
		(*b)[dest.i][dest.j] = nil
		return
	}
	info := PieceInfo{player: piece.pieceInfo().player, IPosn: dest}
	(*b)[dest.i][dest.j] = piece.copy(info)
}

func (b BoardImp) pieces() (pieces []Piece) {
	for _, row := range b {
		for _, piece := range row {
			if piece != nil {
				pieces = append(pieces, piece)
			}
		}
	}
	return pieces
}

func (b BoardImp) deepCopy() Board {
	boardCopy := NewBoardImp()
	for i := range b {
		for j := range b[0] {
			pieceCopy := b[i][j]
			(*boardCopy)[i][j] = pieceCopy
		}
	}
	return boardCopy
}

func (b BoardImp) fen() (s string) {
	for _, row := range b {
		consecBlanks := 0
		for _, piece := range row {
			if piece != nil {
				// add blanks before
				if consecBlanks > 0 {
					s += strconv.Itoa(consecBlanks)
				}
				consecBlanks = 0
				s += piece.String()
			} else {
				consecBlanks++
			}
		}
		if consecBlanks > 0 {
			s += strconv.Itoa(consecBlanks)
		}
		s += "/"
	}
	return s[:len(s)-1] // remove last slash
}
