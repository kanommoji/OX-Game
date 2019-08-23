package ox

import (
	"testing"
)

func Test_Marking_Input_X_Row_0_Column_0_Should_Be_Row_0_Column_0_Is_X(t *testing.T) {
	expected := "x"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.marking(0, 0)
	actual := game.Board[0][0]

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_Marking_Input_O_Row_1_Column_0_Should_Be_Row_1_Column_0_Is_O(t *testing.T) {
	expected := "o"

	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "o")

	game.marking(1, 0)
	actual := game.Board[1][0]

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_Marking_Input_X_Row_0_Column_0_And_O_Row_0_Column_0_Should_Be_Row_0_Column_0_Is_X(t *testing.T) {
	expected := "x"

	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(0, 0)
	game.Play(0, 0)
	actual := game.Board[0][0]

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SwitchTurn_Input_x_Play_Should_Be_Turn_o(t *testing.T) {
	expected := "o"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.switchTurn()
	actual := game.Turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SwitchTurn_Input_o_Play_Should_Be_Turn_x(t *testing.T) {
	expected := "x"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "o")

	game.switchTurn()
	actual := game.Turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckWin_Input_X_Mark_Horizon_First_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.marking(0, 0)
	game.marking(0, 1)
	game.marking(0, 2)
	actual := game.checkWin()

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckWin_Input_X_Mark_Row_0_Column_0_Should_Be_Next_Turn(t *testing.T) {
	expected := "Next Turn"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.marking(0, 0)
	actual := game.checkWin()

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckWin_Input_O_Mark_Horizon_Third_Line_Should_Be_O_Winner(t *testing.T) {
	expected := "o Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "o")

	game.marking(1, 0)
	game.marking(2, 0)
	game.marking(0, 0)
	game.marking(2, 1)
	game.marking(0, 1)
	game.marking(2, 2)
	actual := game.checkWin()

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckFullBoard_Input_Mark_X_And_O_Full_Board_Should_Be_True(t *testing.T) {
	expected := true
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.marking(0, 0)
	game.marking(0, 1)
	game.marking(0, 2)
	game.marking(1, 0)
	game.marking(1, 1)
	game.marking(1, 2)
	game.marking(2, 0)
	game.marking(2, 1)
	game.marking(2, 2)
	actual := game.checkFullBoard()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CheckFullBoard_Input_Empty_Board_Should_Be_False(t *testing.T) {
	expected := false
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	actual := game.checkFullBoard()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CheckFullBoard_Input_Mark_X_Non_Full_Board_Should_Be_False(t *testing.T) {
	expected := false
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.marking(0, 0)
	game.marking(0, 1)
	game.marking(0, 2)
	actual := game.checkFullBoard()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
