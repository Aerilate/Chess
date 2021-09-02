package main

import "fmt"

type ChecksBoard [][]uint

func NewChecksBoard() ChecksBoard {
	board := make([][]uint, BoardSize, BoardSize)

	for i := 0; i < BoardSize; i++ {
		board[i] = make([]uint, BoardSize, BoardSize)
	}
	return board
}

func calcChecksFromBoard(board Board, player int) ChecksBoard {
	threats := make([]IPosn, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != nil && board[i][j].pieceInfo().player != player {
				threats = append(threats, board[i][j].threats(board)...)
			}
		}
	}

	checks := NewChecksBoard()
	for _, posn := range threats {
		if moveInBounds(posn) { // filter
			checks[posn.i][posn.j]++
		}
	}
	return checks
}

func (b *ChecksBoard) squareIsSafe(p IPosn) bool {
	return (*b)[p.i][p.j] == 0
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
		for _, square := range row {
			str += "|"
			if square > 0 {
				str += fmt.Sprintf("%d", square)
			} else {
				str += " "
			}
		}
		str += "|\n"
	}
	addBorder(&str)
	return str
}
