package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MovesList []string

type ViewController struct {
	Board
	activePlayer int
	moves        MovesList
}

func moveInBounds(p Posn) bool {
	return 0 <= p.i && p.i < BoardSize && 0 <= p.j && p.j < BoardSize
}

func (vc *ViewController) move(src Posn, dest Posn) error {
	if !moveInBounds(src) {
		return InvalidMove{"Source coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return InvalidMove{"Destination coordinate " + dest.String() + " is out of range!"}
	} else if src == dest {
		return InvalidMove{"Source and destination coordinates can't be the same!"}
	}

	piece := *vc.at(src)
	if piece == nil { // check piece exists at src
		return InvalidMove{"Coordinate " + src.String() + " has no piece!"}
	} else if piece.player != vc.activePlayer {
		return NotYourPiece{}
	}

	// check piece can move to dest
	err := piece.checkMove(vc, dest)
	if err != nil {
		return err
	}

	*vc.at(dest) = piece
	piece.Posn = dest
	*vc.at(src) = nil
	return nil
}

func (vc *ViewController) loadQueue(moves MovesList) {
	for i := len(moves) - 1; i >= 0; i-- {
		vc.moves = append(vc.moves, moves[i])
	}
}

func (vc *ViewController) nextMove() string {
	l := len(vc.moves)
	if l > 0 {
		nextMove := vc.moves[l-1]
		vc.moves = vc.moves[:l-1]
		return nextMove
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1) // remove \n
	return input
}

func (vc *ViewController) start() {
	vc.Board = *NewBoard()
	vc.activePlayer = Player1

	for {
		fmt.Println(vc.Board)
		fmt.Printf("Player %d's Turn. Enter a move: ", vc.activePlayer)

		for {
			next := vc.nextMove()

			moveRegex, _ := regexp.Compile("m[0-7][0-7][0-7][0-7]")
			if next == "q" || next == "quit" {
				fmt.Println("GAME OVER")
				os.Exit(0)
			} else if moveRegex.MatchString(next) {
				src := Posn{}
				src.i, _ = strconv.Atoi(string(next[1]))
				src.j, _ = strconv.Atoi(string(next[2]))
				dest := Posn{}
				dest.i, _ = strconv.Atoi(string(next[3]))
				dest.j, _ = strconv.Atoi(string(next[4]))

				err := vc.move(src, dest)
				if err == nil { // end turn
					break
				} else {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Println("Unrecognized instruction. Try again: ")
			}
		}
		vc.activePlayer = otherPlayer(vc.activePlayer) // switch players
	}
}
