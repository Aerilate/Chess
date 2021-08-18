package main

type PieceInfo struct {
	player int
	Posn
}

type Posn struct {
	x int
	y int
}

func NewPInfo(player int, x int, y int) PieceInfo {
	p := PieceInfo{player: player, Posn: Posn{x: x, y: y}}
	return p
}
