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

func (game *Game) Play(row, column int) string {
	statusMarking := game.marking(row, column)
	winner := game.checkWin()
	if statusMarking {
		game.switchTurn()
	}
	return winner
}

func (game Game) checkBoardPositionEmpty(row, column int) bool {
	return game.Board[row][column] == ""
}

func (game *Game) marking(row, column int) bool {
	if !game.checkBoardPositionEmpty(row, column) {
		return false
	}
	game.Board[row][column] = game.turn
	return true
}

func (game Game) checkWin() string {
	if game.checkWinHorizontal() ||
		game.checkWinVertical() ||
		game.checkWinDiagonal() {
		return game.turn + " Winner"
	}
	if game.checkFullBoard() {
		return "Tie"
	}
	return "Next Turn"
}

func (game Game) checkWinDiagonal() bool {
	if game.Board[0][0] == game.turn &&
		game.Board[1][1] == game.turn &&
		game.Board[2][2] == game.turn {
		return true
	}
	if game.Board[2][0] == game.turn &&
		game.Board[1][1] == game.turn &&
		game.Board[0][2] == game.turn {
		return true
	}
	return false
}
func (game Game) checkWinVertical() bool {
	for index, _ := range game.Board {
		if game.Board[0][index] == game.turn &&
			game.Board[1][index] == game.turn &&
			game.Board[2][index] == game.turn {
			return true
		}
	}
	return false
}

func (game Game) checkWinHorizontal() bool {
	for index, _ := range game.Board {
		if game.Board[index][0] == game.turn &&
			game.Board[index][1] == game.turn &&
			game.Board[index][2] == game.turn {
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
