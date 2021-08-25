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
