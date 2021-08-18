package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := NewBoard()
	reader := bufio.NewReader(os.Stdin)
	player := 1

	for {
		fmt.Println(b)
		fmt.Printf("Player %d's Turn. Enter a move: ", player)

		for {
			in, _ := reader.ReadString('\n')
			in = strings.Replace(in, "\n", "", -1) // remove \n

			switch in {
			case "q", "quit":
				fmt.Println("GAME OVER")
				os.Exit(0)
			case "m":
				// valid move
			default:
				fmt.Print("Invalid move. Try again: ")
			}
		}
		player = 3 - player
	}
}
