package main

import (
	"fmt"
)

// IPosn slice coordinates representation
type IPosn struct {
	i int
	j int
}

func (p IPosn) toStdPosn() StdPosn {
	file := p.j + 'a'
	rank := BoardSize - p.i
	return StdPosn{rune(file), rank}
}

func (p IPosn) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

func (p IPosn) add(other IPosn) IPosn {
	return IPosn{p.i + other.i, p.j + other.j}
}

func filter(posns []IPosn, fn func(IPosn) bool) (result []IPosn) {
	for _, posn := range posns {
		if fn(posn) {
			result = append(result, posn)
		}
	}
	return result
}

// StdPosn standard Chess notation
type StdPosn struct {
	file rune
	rank int
}

func (p StdPosn) toIPosn() IPosn {
	i := BoardSize - p.rank
	j := p.file - 'a'
	return IPosn{i, int(j)}
}

func (p StdPosn) String() string {
	return fmt.Sprintf("(%c,%d)", p.file, p.rank)
}

type BadStdPosn struct{}

func (e BadStdPosn) Error() string {
	return "Can't convert string to a coordinate!"
}
