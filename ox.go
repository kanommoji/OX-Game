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
	if game.Board[0][0] == player.symbol && game.Board[0][1] == player.symbol && game.Board[0][2] == player.symbol {
		return player.symbol + " Winner"
	}
	if game.Board[1][0] == player.symbol && game.Board[1][1] == player.symbol && game.Board[1][2] == player.symbol {
		return player.symbol + " Winner"
	}
	if game.Board[2][0] == player.symbol && game.Board[2][1] == player.symbol && game.Board[2][2] == player.symbol {
		return player.symbol + " Winner"
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
