package main

const Player1 = 1
const Player2 = 2

// maps 1->6 and 2->1
func pawnHomeRow(player int) int {
	return (BoardSize/2+1)*(2-player) + 1
}

// maps 1->-1 and 2->1
func moveDirection(player int) int {
	return player*2 - 3
}

func otherPlayer(player int) int {
	return 3 - player
}
