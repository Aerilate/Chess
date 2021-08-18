package main

type King struct {
	posn
}

func (k King) String() string {
	return "k"
}
