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
	winner := game.checkWin(player)
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, row, column int) {
	game.Board[row][column] = player.symbol
}

func (game Game) checkWin(player Player) string {
	for index, _ := range game.Board {
		if game.Board[index][0] == player.symbol && game.Board[index][1] == player.symbol && game.Board[index][2] == player.symbol {
			return player.symbol + " Winner"
		}
	}
	return "Next Turn"
}

func (game *Game) switchTurn() {
	if game.turn == "x" {
		game.turn = "o"
	} else {
		game.turn = "x"
	}
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
