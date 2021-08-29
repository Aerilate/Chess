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
	activePlayer int // 1 or 2
	moveHistory  []Move
	checks       [3]ChecksBoard
}

func NewGameState() *GameState {
	newGame := &GameState{Board: *NewBoard(), activePlayer: 1, checks: [3]ChecksBoard{*NewChecksBoard(), *NewChecksBoard(), *NewChecksBoard()}}
	newGame.updateChecks()
	return newGame
}

func (game *GameState) updateChecks() {
	threats := make([]IPosn, 0)
	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[0]); j++ {
			if game.Board[i][j] != nil && game.Board[i][j].player != game.activePlayer {
				threats = append(threats, game.Board[i][j].threats(game.Board)...)
			}
		}
	}

	for _, posn := range threats {
		if moveInBounds(posn) { // filter
			game.checks[game.activePlayer][posn.i][posn.j]++
		}
	}
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
	} else if piece.player != game.activePlayer {
		return NotYourPiece{}
	} else if *game.at(dest) != nil && piece.player == (*game.at(dest)).player { // check if dest occupied by own piece
		return InvalidMove{"Coordinate " + dest.String() + " is occupied by your own piece!"}
	}

	// check piece can move to dest
	err := piece.checkMove(game.Board, game.checks[game.activePlayer], dest)
	if err != nil {
		return err
	}

	// all is good!
	*game.at(dest) = piece
	piece.IPosn = dest
	*game.at(src) = nil
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
	return game.Board.String() + "\n" + game.checks[game.activePlayer].String()
}
