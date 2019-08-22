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

func (game Game) Play(player Player, row, column int) string {
	game.marking(player, row, column)
	winner := game.checkWin()
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, row, column int) {
	game.Board[row][column] = player.symbol
}

func (game Game) checkWin() string {
	return "x Winner"
}

func (game Game) switchTurn() {
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
