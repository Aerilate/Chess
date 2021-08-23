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
		return InvalidMove{"Source coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return InvalidMove{"Destination coordinate " + dest.String() + " is out of range!"}
	} else if src.equals(dest) {
		return InvalidMove{"Source and destination coordinates can't be the same!"}
	}

	piece := *v.at(src)
	if piece == nil {
		return InvalidMove{"Coordinate " + src.String() + " has no piece!"}
	}
	err := piece.checkMove(&v, dest)
	if err != nil {
		return err
	}

	*v.at(dest) = *v.at(src)
	*v.at(src) = nil

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
				src := Posn{}
				src.x, _ = strconv.Atoi(string(in[1]))
				src.y, _ = strconv.Atoi(string(in[2]))
				dest := Posn{}
				dest.x, _ = strconv.Atoi(string(in[3]))
				dest.y, _ = strconv.Atoi(string(in[4]))

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
