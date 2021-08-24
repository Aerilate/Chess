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
	return 0 <= p.i && p.i < BoardSize && 0 <= p.j && p.j < BoardSize
}

func (vc ViewController) move(src Posn, dest Posn) error {
	if !moveInBounds(src) {
		return InvalidMove{"Source coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return InvalidMove{"Destination coordinate " + dest.String() + " is out of range!"}
	} else if src.equals(dest) {
		return InvalidMove{"Source and destination coordinates can't be the same!"}
	}

	piece := *vc.at(src)
	if piece == nil { // check piece exists at src
		return InvalidMove{"Coordinate " + src.String() + " has no piece!"}
	} else if piece.player != vc.activePlayer {
		return NotYourPiece{}
	} else if *vc.at(dest) != nil && (*vc.at(dest)).player == vc.activePlayer { // check if dest occupied by own piece
		return InvalidMove{"Coordinate " + dest.String() + " is occupied by your piece!"}
	}

	// check piece can move to dest
	err := piece.checkMove(&vc, dest)
	if err != nil {
		return err
	}

	*vc.at(dest) = *vc.at(src)
	*vc.at(src) = nil
	return nil
}

func (vc ViewController) start() {
	vc.Board = *NewBoard()
	vc.activePlayer = Player1
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(vc.Board)
		fmt.Printf("Player %d's Turn. Enter a move: ", vc.activePlayer)

		for {
			in, _ := reader.ReadString('\n')
			in = strings.Replace(in, "\n", "", -1) // remove \n

			moveRegex, _ := regexp.Compile("m[0-7][0-7][0-7][0-7]")
			if in == "q" || in == "quit" {
				fmt.Println("GAME OVER")
				os.Exit(0)
			} else if moveRegex.MatchString(in) {
				src := Posn{}
				src.i, _ = strconv.Atoi(string(in[1]))
				src.j, _ = strconv.Atoi(string(in[2]))
				dest := Posn{}
				dest.i, _ = strconv.Atoi(string(in[3]))
				dest.j, _ = strconv.Atoi(string(in[4]))

				err := vc.move(src, dest)
				if err == nil { // end turn
					break
				} else {
					fmt.Print(err.Error())
				}
			} else {
				fmt.Print("Unrecognized instruction. Try again: ")
			}
		}
		vc.activePlayer = Player1 + Player2 - vc.activePlayer // switch players
	}
}
