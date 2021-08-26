package main

import (
	"bufio"
	"os"
)

func gameReader(fileName string) (slice []string, err error) {
	file, err := os.Open(fileName)
	defer file.Close()
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
