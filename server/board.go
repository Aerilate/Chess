package main

const BoardSize = 8

type Board [][]Piece

func NewBoard() *Board {
	board := make([][]Piece, BoardSize, BoardSize)
	for i := 0; i < BoardSize; i++ {
		board[i] = make([]Piece, BoardSize, BoardSize)
	}
	return (*Board)(&board)
}

func moveInBounds(p IPosn) bool {
	return 0 <= p.i && p.i < BoardSize && 0 <= p.j && p.j < BoardSize
}

func (b *Board) at(p IPosn) *Piece {
	return &(*b)[p.i][p.j]
}

func (b Board) String() string {
	addBorder := func(str *string) {
		*str += "O"
		for i := 0; i < BoardSize*2-1; i++ {
			*str += "="
		}
		*str += "O\n"
	}

	str := ""
	addBorder(&str)
	for _, row := range b {
		for _, piece := range row {
			str += "|"
			if piece == nil {
				str += " "
			} else {
				str += piece.String()
			}
		}
		str += "|\n"
	}
	addBorder(&str)
	return str
}

func (b Board) kingUnderCheck(player int) bool {
	kingPosn := IPosn{}
	for _, row := range b {
		for _, piece := range row {
			switch piece.(type) {
			case *King:
				kingPosn = piece.pieceInfo().IPosn
			}
		}
	}
	checks := calcChecksFromBoard(b, player)
	return !checks.squareIsSafe(kingPosn)
}

func (b Board) shallowCopy() Board {
	boardCopy := *NewBoard()
	for i := range b {
		for j := range b[0] {
			boardCopy[i][j] = b[i][j]
		}
	}
	return boardCopy
}
