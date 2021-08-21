package main

import (
	"fmt"
)

type Posn struct {
	x int
	y int
}

func (p Posn) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
