package chess

const Player1 = 1
const Player2 = 2

type Player struct {
	int
	king            *King
	pieces          map[Piece]bool
	attackedSquares ChecksBoard
}

func (player *Player) NewPiece(pieceType string, posn IPosn) (p Piece) {
	info := PieceInfo{player.int, posn}
	switch pieceType {
	case pawn:
		p = &Pawn{info}
	case knight:
		p = &Knight{info}
	case bishop:
		p = &Bishop{info}
	case rook:
		p = &Rook{info}
	case queen:
		p = &Queen{info}
	case king:
		player.king = &King{info}
		p = player.king
	}

	if player.pieces == nil {
		player.pieces = make(map[Piece]bool)
	}
	player.pieces[p] = true
	return p
}

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
