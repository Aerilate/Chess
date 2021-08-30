package main

type Gameable interface {
	move(src IPosn, dest IPosn) error
	getActivePlayer() int
	String() string
}

type Move struct {
	src  IPosn
	dest IPosn
}

type GameState struct {
	Board
	prevBoard Board

	players      [3]Player
	activePlayer int // 1 or 2

	moveHistory []Move
}

func NewGameState() *GameState {
	newGame := &GameState{Board: *NewBoard(), activePlayer: Player2}
	newGame.players[1] = Player{int: Player1}
	newGame.players[2] = Player{int: Player2}
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

func (game *GameState) undo() {
	movesMade := len(game.moveHistory)
	if movesMade == 0 {
		return
	}
	game.Board = game.prevBoard
	game.moveHistory = game.moveHistory[:movesMade-1]
}

func calcThreats(board Board, player int) []IPosn {
	threats := make([]IPosn, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != nil && board[i][j].pieceInfo().player != player {
				threats = append(threats, board[i][j].threats(board)...)
			}
		}
	}
	return threats
}

func (game *GameState) updateChecks() {
	game.players[game.activePlayer].attackedSquares = *NewChecksBoard() // clear board
	threats := calcThreats(game.Board, game.activePlayer)
	for _, posn := range threats {
		if moveInBounds(posn) { // filter
			game.players[game.activePlayer].attackedSquares[posn.i][posn.j]++
		}
	}
}

func (game *GameState) rollbackMove() {
	game.Board = game.prevBoard
	for i := range game.Board {
		for j := range game.Board[0] {
			posn := IPosn{i, j}
			if *game.Board.at(posn) != nil {
				(*game.Board.at(posn)).updatePosn(posn)
			}
		}
	}
	game.updateChecks()
}

func (game *GameState) move(src IPosn, dest IPosn) error {
	if !moveInBounds(src) {
		return InvalidMove{"Source coordinate " + src.String() + " is out of range!"}
	} else if !moveInBounds(dest) {
		return InvalidMove{"Destination coordinate " + dest.String() + " is out of range!"}
	} else if src == dest {
		return InvalidMove{"Source and destination coordinates can't be the same!"}
	}

	piece := *game.at(src)
	if piece == nil { // check piece exists at src
		return InvalidMove{"Coordinate " + src.String() + " has no piece!"}
	} else if piece.pieceInfo().player != game.activePlayer {
		return NotYourPiece{}
	} else if *game.at(dest) != nil && piece.pieceInfo().player == (*game.at(dest)).pieceInfo().player { // check if dest occupied by own piece
		return InvalidMove{"Coordinate " + dest.String() + " is occupied by your own piece!"}
	}

	// check piece can move to dest
	err := piece.checkMove(game.Board, game.players[game.activePlayer].attackedSquares, dest)
	if err != nil {
		return err
	}

	game.prevBoard = game.Board.shallowCopy()
	*game.at(dest) = piece
	piece.updatePosn(dest)
	*game.at(src) = nil

	game.updateChecks()
	if game.players[game.activePlayer].king.underCheck(game.players[game.activePlayer].attackedSquares) {
		game.rollbackMove()
		return InvalidMove{"King would be under check!"}
	}

	game.moveHistory = append(game.moveHistory, Move{src, dest})
	game.endTurn()
	return nil
}

func (game *GameState) endTurn() {
	game.activePlayer = otherPlayer(game.activePlayer)
	game.updateChecks()
}

func (game GameState) getActivePlayer() int {
	return game.activePlayer
}

func (game GameState) String() string {
	return game.Board.String() + "\n" + game.players[otherPlayer(game.activePlayer)].attackedSquares.String()
}
