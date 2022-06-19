package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readMove() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", fmt.Errorf("failed to read from scanner")
	}
	return scanner.Text(), scanner.Err()
}

func main() {
	game := NewGame()
	for !game.IsOver() {
		print(game)
		input, err := readMove()
		if err != nil {
			fmt.Println(err)
		}

		rank, _ := strconv.ParseInt(string(input[1]), 10, 32)
		src := StdPosn{file: rune(input[0]), rank: int(rank)}
		rank, _ = strconv.ParseInt(string(input[3]), 10, 32)
		dest := StdPosn{file: rune(input[2]), rank: int(rank)}
		move := Move{src, dest}
		fmt.Printf("%+v\n", move)
		game.Move(move)
	}
}
