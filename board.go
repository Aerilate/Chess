package main

const BoardSize = 8

type Board [][]*Piece

func NewBoard() *Board {
	board := Board{}
	board = make([][]*Piece, BoardSize, BoardSize)

	for i := 0; i < BoardSize; i++ {
		board[i] = make([]*Piece, BoardSize, BoardSize)
		for j := 0; j < BoardSize; j++ {
			slot := &board[i][j]
			player := Player1
			if i == 0 || i == 1 {
				player = Player2
			}

			if i == 0 || i == 7 {
				posn := IPosn{i, j}
				switch j {
				case 0, 7:
					*slot = &Piece{rook, player, posn}
				case 1, 6:
					*slot = &Piece{knight, player, posn}
				case 2, 5:
					*slot = &Piece{bishop, player, posn}
				case 3:
					*slot = &Piece{queen, player, posn}
				case 4:
					*slot = &Piece{king, player, posn}
				}
			} else if i == 1 || i == 6 {
				posn := IPosn{i, j}
				*slot = &Piece{pawn, player, posn}
			}
		}
	}
	return &board
}

func (b *Board) at(p IPosn) **Piece {
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
