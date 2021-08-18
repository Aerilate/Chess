package main

type Queen struct {
	posn
}

func (q Queen) String() string {
	return "q"
}
