package main

import (
	"fmt"
)

// IPosn slice coordinates representation
type IPosn struct {
	i int
	j int
}

func (p *IPosn) add(other IPosn) {
	p.i += other.i
	p.j += other.j
}

func (p IPosn) toStdPosn() StdPosn {
	file := p.j + 'a'
	rank := main.BoardSize - p.i
	return StdPosn{rune(file), rank}
}

func (p IPosn) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

// StdPosn standard Chess notation
type StdPosn struct {
	file rune
	rank int
}

func (p StdPosn) toIPosn() IPosn {
	i := main.BoardSize - p.rank
	j := p.file - 'a'
	return IPosn{i, int(j)}
}

func (p StdPosn) String() string {
	return fmt.Sprintf("%c%d", p.file, p.rank)
}
