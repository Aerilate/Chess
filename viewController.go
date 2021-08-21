package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Player1 = 1
const Player2 = 2

type ViewController struct {
	Board
	activePlayer int
}

func moveInBounds(p Posn) bool {
	return 0 <= p.x && p.x < BoardSize && 0 <= p.y && p.y < BoardSize
}

func (v ViewController) move(src Posn, dest Posn) error {
	if !moveInBounds(src) {
		return InvalidMove{"Coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return InvalidMove{"Coordinate " + dest.String() + " is out of range!"}
	}

	if v.board[src.x][src.y] == nil {
		return InvalidMove{"Coordinate " + src.String() + " has no piece!"}
	}

	v.board[dest.x][dest.y] = v.board[src.x][src.y]
	v.board[src.x][src.y] = nil

	return nil
}

func (v ViewController) start() {
	v.Board = *NewBoard()
	v.activePlayer = Player1
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(v.Board)
		fmt.Printf("Player %d's Turn. Enter a move: ", v.activePlayer)

		for {
			in, _ := reader.ReadString('\n')
			in = strings.Replace(in, "\n", "", -1) // remove \n

			moveRegex, _ := regexp.Compile("m[0-7][0-7][0-7][0-7]")
			if in == "q" || in == "quit" {
				fmt.Println("GAME OVER")
				os.Exit(0)
			} else if moveRegex.MatchString(in) {
				a, _ := strconv.Atoi(string(in[1]))
				b, _ := strconv.Atoi(string(in[2]))
				src := Posn{a, b}
				a, _ = strconv.Atoi(string(in[3]))
				b, _ = strconv.Atoi(string(in[4]))
				dest := Posn{a, b}

				err := v.move(src, dest)
				if err == nil { // end turn
					break
				} else {
					fmt.Print(err.Error())
				}
			} else {
				fmt.Print("Unrecognized instruction. Try again: ")
			}
		}
		v.activePlayer = Player1 + Player2 - v.activePlayer // switch players
	}
}

type InvalidMove struct {
	s string
}

func (e InvalidMove) Error() string {
	return "Invalid move, " + e.s + " Try again: "
}
