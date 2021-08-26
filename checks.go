package main

import "fmt"

type ChecksBoard [][]uint

func NewChecksBoard() *ChecksBoard {
	board := make([][]uint, BoardSize, BoardSize)

	for i := 0; i < BoardSize; i++ {
		board[i] = make([]uint, BoardSize, BoardSize)
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
