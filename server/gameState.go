package main

import "strconv"

const Player1 = 1
const Player2 = 2

type GameImp struct {
	Board
	activePlayer int
	gameOver     bool
}

func NewGameState() *GameImp {
	newGame := &GameImp{Board: *NewBoard()}
	newGame.activePlayer = Player1
	newGame.setupPieces()
	return newGame
}

func (game *GameImp) setupPieces() {
	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[0]); j++ {
			slot := &game.Board[i][j]
			player := Player1
			if i == 0 || i == 1 {
				player = Player2
			}
			info := PieceInfo{player: player, IPosn: IPosn{i, j}}

			if i == 0 || i == 7 {
				switch j {
				case 0, 7:
					*slot = NewPiece(rook, info)
				case 1, 6:
					*slot = NewPiece(knight, info)
				case 2, 5:
					*slot = NewPiece(bishop, info)
				case 3:
					*slot = NewPiece(queen, info)
				case 4:
					*slot = NewPiece(king, info)
				}
			} else if i == 1 || i == 6 {
				*slot = NewPiece(pawn, info)
			}
		}
	}
}

func (game *GameImp) validMoves() (validMoves map[string][]string) {
	validMoves = make(map[string][]string)
	movesLeft := false

	for _, row := range game.Board {
		for _, piece := range row {
			if piece.pieceInfo().player == game.activePlayer {
				iMoves := piece.validMoves(game.Board)

				// convert to StdPosn
				stdMoves := make([]string, len(iMoves))
				for i, move := range iMoves {
					stdMoves[i] = move.toStdPosn().String()
				}

				validMoves[piece.pieceInfo().IPosn.toStdPosn().String()] = stdMoves
				movesLeft = movesLeft || len(iMoves) > 0
			}
		}
	}

	game.gameOver = !movesLeft
	return validMoves
}

func (game *GameImp) move(move Move) {
	src := move.src.toIPosn()
	dest := move.dest.toIPosn()

	// update board
	piece := *game.at(src)
	*game.at(dest) = piece
	piece.updatePosn(dest)
	*game.at(src) = nil

	game.activePlayer = 3 - game.activePlayer // switch players
}

func (game GameImp) ActivePlayer() int {
	return game.activePlayer
}

func (game *GameImp) isOver() bool {
	return game.gameOver
}

func (game *GameImp) fen() (s string) {
	for _, row := range game.Board {
		consecBlanks := 0
		for _, piece := range row {
			if piece != nil {
				// add blanks before
				if consecBlanks > 0 {
					s += strconv.Itoa(consecBlanks)
				}
				consecBlanks = 0
				s += piece.String()
			} else {
				consecBlanks++
			}
		}
		if consecBlanks > 0 {
			s += strconv.Itoa(consecBlanks)
		}
		s += "/"
	}
	return s[:len(s)-1] // remove last slash
}
