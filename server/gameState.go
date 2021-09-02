package main

type GameState struct {
	Board
	prevBoard Board

	players      [3]Player
	activePlayer *Player

	moveHistory []Move
	gameOver    bool
}

func (game *GameState) validMoves() (validMoves map[string][]string) {
	validMoves = make(map[string][]string)
	movesLeft := false
	for piece := range game.activePlayer.pieces {
		iMoves := piece.validMoves(game.Board)
		// convert to StdPosn
		stdMoves := make([]string, len(iMoves))
		for i, move := range iMoves {
			stdMoves[i] = move.toStdPosn().String()
		}

		validMoves[piece.pieceInfo().IPosn.toStdPosn().String()] = stdMoves
		movesLeft = movesLeft || len(iMoves) > 0
	}
	game.gameOver = !movesLeft
	return validMoves
}

func (game *GameState) isOver() bool {
	return game.gameOver
}

func (game *GameState) lastMove() Move {
	if len(game.moveHistory) == 0 {
		return Move{}
	}
	return game.moveHistory[len(game.moveHistory)-1]
}

func NewGameState() *GameState {
	newGame := &GameState{Board: *NewBoard()}
	newGame.players[1] = Player{int: Player1}
	newGame.players[2] = Player{int: Player2}
	newGame.activePlayer = &newGame.players[2]
	newGame.setupPieces()
	newGame.endTurn()
	return newGame
}

func (game *GameState) setupPieces() {
	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[0]); j++ {
			slot := &game.Board[i][j]
			player := &game.players[Player1]
			if i == 0 || i == 1 {
				player = &game.players[Player2]
			}

			if i == 0 || i == 7 {
				posn := IPosn{i, j}
				switch j {
				case 0, 7:
					*slot = player.NewPiece(rook, posn)
				case 1, 6:
					*slot = player.NewPiece(knight, posn)
				case 2, 5:
					*slot = player.NewPiece(bishop, posn)
				case 3:
					*slot = player.NewPiece(queen, posn)
				case 4:
					*slot = player.NewPiece(king, posn)
				}
			} else if i == 1 || i == 6 {
				posn := IPosn{i, j}
				*slot = player.NewPiece(pawn, posn)
			}
		}
	}
}

func (game *GameState) move(move Move) {
	src := move.src.toIPosn()
	dest := move.dest.toIPosn()

	// update board
	piece := *game.at(src)
	*game.at(dest) = piece
	piece.updatePosn(dest)
	*game.at(src) = nil

	game.moveHistory = append(game.moveHistory, Move{move.src, move.dest})
	game.endTurn()
}

func (game GameState) inactivePlayer() *Player {
	otherPlayerId := otherPlayer(game.activePlayer.int)
	return &game.players[otherPlayerId]
}

func (game *GameState) endTurn() {
	game.activePlayer = game.inactivePlayer()
}

func (game GameState) getActivePlayer() int {
	return game.activePlayer.int
}

func (game GameState) String() string {
	return game.Board.String()
}
