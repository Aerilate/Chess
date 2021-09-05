package main

const (
	pawn   = "p"
	knight = "n"
	bishop = "b"
	rook   = "r"
	queen  = "q"
	king   = "k"
)

type Piece interface {
	threats(board Board) (threats []IPosn)
	validDests(board Board) (dests []IPosn)
	pieceInfo() PieceInfo
	copy(info PieceInfo) Piece
	String() string
}

type PieceInfo struct {
	player int
	IPosn
}

func NewPiece(pieceType string, info PieceInfo) (p Piece) {
	switch pieceType {
	case pawn:
		p = Pawn{info}
	case knight:
		p = Knight{info}
	case bishop:
		p = Bishop{info}
	case rook:
		p = Rook{info}
	case queen:
		p = Queen{info}
	case king:
		p = King{info}
	}
	return p
}
