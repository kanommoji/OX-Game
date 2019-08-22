package ox

type Game struct {
	Board  [3][3]string
	player []Player
	turn   string
}

type Player struct {
	name   string
	symbol string
	score  int
}

func (game *Game) Play(player Player, row, column int) string {
	game.marking(player, row, column)
	winner := game.checkWin(player.symbol)
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, row, column int) {
	game.Board[row][column] = player.symbol
}

func (game Game) checkWin(playerSymbol string) string {
	if game.checkWinHorizontal(playerSymbol) ||
		game.checkWinVertical(playerSymbol) ||
		game.checkWinDiagonal(playerSymbol) {
		return playerSymbol + " Winner"
	}
	if game.checkFullBoard() {
		return "Tie"
	}
	return "Next Turn"
}

func (game Game) checkWinDiagonal(playerSymbol string) bool {
	if game.Board[0][0] == playerSymbol &&
		game.Board[1][1] == playerSymbol &&
		game.Board[2][2] == playerSymbol {
		return true
	}
	if game.Board[2][0] == playerSymbol &&
		game.Board[1][1] == playerSymbol &&
		game.Board[0][2] == playerSymbol {
		return true
	}
	return false
}
func (game Game) checkWinVertical(playerSymbol string) bool {
	for index, _ := range game.Board {
		if game.Board[0][index] == playerSymbol &&
			game.Board[1][index] == playerSymbol &&
			game.Board[2][index] == playerSymbol {
			return true
		}
	}
	return false
}

func (game Game) checkWinHorizontal(playerSymbol string) bool {
	for index, _ := range game.Board {
		if game.Board[index][0] == playerSymbol &&
			game.Board[index][1] == playerSymbol &&
			game.Board[index][2] == playerSymbol {
			return true
		}
	}
	return false
}

func (game *Game) switchTurn() {
	if game.turn == "x" {
		game.turn = "o"
	} else {
		game.turn = "x"
	}
}

func (game Game) checkFullBoard() bool {
	for _, board2D := range game.Board {
		for _, board := range board2D {
			if board == "" {
				return false
			}
		}
	}
	return true
}

func NewGame(player1, player2 Player, turn string) Game {
	return Game{
		Board:  [3][3]string{},
		player: []Player{player1, player2},
		turn:   turn,
	}
}

func NewPlayer(name string, symbol string) Player {
	return Player{
		name:   name,
		score:  0,
		symbol: symbol,
	}
}
