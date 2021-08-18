package main

type Bishop struct {
	posn
}

func (b Bishop) String() string {
	return "b"
}
