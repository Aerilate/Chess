package main

const BoardSize = 8

type Board struct {
	board [BoardSize][BoardSize]piece
}

func New() *Board {
	b := Board{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			slot := &b.board[i][j]
			if i == 0 {
				if j == 0 || j == 7 {
					*slot = Rook{}
				} else if j == 1 || j == 6 {
					*slot = Knight{}
				} else if j == 2 || j == 5 {
					*slot = Bishop{}
				} else if j == 3 {
					*slot = Queen{}
				} else if j == 4 {
					*slot = King{}
				}
			} else if i == 1 {
				*slot = Pawn{}
			}
		}
	}
	return &b
}

func addBorder(str *string) {
	*str += "O"
	for i := 0; i < BoardSize*2-1; i++ {
		*str += "="
	}
	*str += "O\n"
}

func (b Board) String() string {
	str := ""
	addBorder(&str)
	for _, row := range b.board {
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
