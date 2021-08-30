package chess

import (
	"fmt"
	"regexp"
	"strconv"
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

// StdPosn standard Chess notation
type StdPosn struct {
	file rune
	rank int
}

func toStdPosn(s string) (*StdPosn, error) {
	moveRegex, _ := regexp.Compile("[a-h][0-7]")
	if !moveRegex.MatchString(s) {
		return nil, BadStdPosn{}
	}

	file := rune(s[0])
	rank, _ := strconv.Atoi(string(s[1]))
	return &StdPosn{file, rank}, nil
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
