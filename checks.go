package main

type ChecksBoard [][]bool

func NewChecksBoard() *ChecksBoard {
	board := make([][]bool, BoardSize, BoardSize)

	for i := 0; i < BoardSize; i++ {
		board[i] = make([]bool, BoardSize, BoardSize)
	}
	return (*ChecksBoard)(&board)
}

func (b ChecksBoard) String() string {
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
			if piece {
				str += "1"
			} else {
				str += " "
			}
		}
		str += "|\n"
	}
	addBorder(&str)
	return str
}
