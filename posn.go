package main

import (
	"fmt"
)

type Posn struct {
	i int
	j int
}

func (p Posn) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

func (p Posn) equals(other Posn) bool {
	return p.i == other.i && p.j == other.j
}
