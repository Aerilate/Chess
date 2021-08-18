package main

const BoardSize = 8
const Player1 = 1
const Player2 = 2

type Board struct {
	board [BoardSize][BoardSize]piece
}

func NewBoard() *Board {
	b := Board{}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			slot := &b.board[i][j]
			player := Player1
			if i == 0 || i == 1 {
				player = Player2
			}

			if i == 0 || i == 7 {
				pInfo := NewPInfo(player, i, j)
				switch j {
				case 0, 7:
					*slot = NewRook(pInfo)
				case 1, 6:
					*slot = NewKnight(pInfo)
				case 2, 5:
					*slot = NewBishop(pInfo)
				case 3:
					*slot = NewQueen(pInfo)
				case 4:
					*slot = NewKing(pInfo)
				}
			} else if i == 1 || i == 6 {
				pInfo := NewPInfo(player, i, j)
				*slot = NewPawn(pInfo)
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
