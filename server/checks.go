package main

import "fmt"

type ChecksBoard [][]uint

func kingUnderCheck(board Board, player int) bool {
	kingPosn := IPosn{}
	for _, piece := range board.pieces() {
		if piece.pieceInfo().player != player {
			continue
		}
		if _, ok := piece.(King); ok {
			kingPosn = piece.pieceInfo().IPosn
		}
	}
	checks := calcChecksFromBoard(board, player)
	return !checks.squareIsSafe(kingPosn)
}

func calcChecksFromBoard(board Board, player int) (checks ChecksBoard) {
	threats := make([]IPosn, 0)
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			posn := IPosn{i, j}
			if board.squareIsEmpty(posn) {
				continue
			}

			piece := board.at(posn)
			if piece.pieceInfo().player != player {
				threats = append(threats, piece.threats(board)...)
			}
		}
	}

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
