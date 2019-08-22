package ox

import (
	"testing"
)

func Test_Marking_Input_X_Row_0_Column_0_Should_Be_Row_0_Column_0_Is_X(t *testing.T) {
	expected := "x"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)

	actual := game.Board[0][0]

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_Marking_Input_O_Row_1_Column_0_Should_Be_Row_1_Column_0_Is_O(t *testing.T) {
	expected := "o"

	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player2, 1, 0)

	actual := game.Board[1][0]

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SwitchTurn_Input_x_Play_Should_Be_Turn_o(t *testing.T) {
	expected := "o"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "x")

	game.switchTurn()
	actual := game.turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SwitchTurn_Input_o_Play_Should_Be_Turn_x(t *testing.T) {
	expected := "x"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "o")

	game.switchTurn()
	actual := game.turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckWin_Input_X_Mark_Horizon_First_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	game.marking(player1, 0, 1)
	game.marking(player1, 0, 2)

	actual := game.checkWin(player1)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckWin_Input_X_Mark_Row_0_Column_0_Should_Be_Next_Turn(t *testing.T) {
	expected := "Next Turn"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	actual := game.checkWin(player1)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
