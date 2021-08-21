package main

import (
	"strings"
)

type PieceInfo struct {
	player int
	Posn
}

func NewPInfo(player int, x int, y int) PieceInfo {
	p := PieceInfo{player: player, Posn: Posn{x: x, y: y}}
	return p
}

func (p *PieceInfo) formatStr(s string) string {
	if p.player == Player1 {
		return strings.ToUpper(s)
	}
	return s
}
