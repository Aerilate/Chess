package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const Player1 = 1
const Player2 = 2

type ViewController struct {
	board        Board
	activePlayer int
}

func moveInBounds(p Posn) bool {
	return 0 <= p.x && p.x < BoardSize && 0 <= p.y && p.y < BoardSize
}

func (v ViewController) move(src Posn, dest Posn) error {
	if !moveInBounds(src) {
		return invalidMove{"Coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return invalidMove{"Coordinate " + dest.String() + " is out of range!"}
	}

	return nil
}

func (v ViewController) start() {
	v.board = *NewBoard()
	v.activePlayer = Player1
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(v.board)
		fmt.Printf("Player %d's Turn. Enter a move: ", v.activePlayer)

		for {
			in, _ := reader.ReadString('\n')
			in = strings.Replace(in, "\n", "", -1) // remove \n

			moveRegex, _ := regexp.Compile("m[0-7][0-7][0-7][0-7]")
			if in == "q" || in == "quit" {
				fmt.Println("GAME OVER")
				os.Exit(0)
			} else if moveRegex.MatchString(in) {
				src := Posn{int(in[1]), int(in[2])}
				dest := Posn{int(in[3]), int(in[4])}

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

type invalidMove struct {
	s string
}

func (e invalidMove) Error() string {
	return "Invalid move, " + e.s + " Try again: "
}
