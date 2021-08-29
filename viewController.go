package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ViewController struct {
	Gameable
}

func readGameFile(fileName string) (slice []string, err error) {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read file")
		}
	}(file)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return slice, err
	}
	return slice, err
}

func (vc *ViewController) loadSavedGame(fileName string) error {
	file, err := readGameFile(fileName)
	if err != nil {
		return err
	}

	for _, line := range file {
		err = vc.move(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vc *ViewController) readNextMove() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1) // remove \n
	return input
}

func (vc *ViewController) move(s string) (err error) {
	src, err := toStdPosn(s[0:2])
	if err != nil {
		return err
	}
	dest, err := toStdPosn(s[2:4])
	if err != nil {
		return err
	}

	err = vc.Gameable.move(src.toIPosn(), dest.toIPosn())
	return err
}

func (vc *ViewController) start() {
	for {
		fmt.Println(vc)
		fmt.Printf("Player %d's Turn. Enter a move: ", vc.getActivePlayer())

		for {
			nextMove := vc.readNextMove()
			if nextMove == "q" || nextMove == "quit" {
				fmt.Println("GAME OVER")
				os.Exit(0)
			} else if moveRegex, _ := regexp.Compile("^[a-h][0-7][a-h][0-7]$"); moveRegex.MatchString(nextMove) {
				err := vc.move(nextMove)
				if err == nil {
					break
				}
				fmt.Println(err.Error())
			} else {
				fmt.Println("Unrecognized instruction. Try again: ")
			}
		}
	}
}
