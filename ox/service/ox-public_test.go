package ox

import (
	"testing"
)

func Test_X_Win_By_Horizon_First_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(0, 0)
	game.Play(1, 1)
	game.Play(0, 1)
	game.Play(1, 0)
	winner := game.Play(0, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Horizon_Second_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(1, 0)
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(0, 0)
	winner := game.Play(1, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Horizon_Third_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(2, 0)
	game.Play(0, 1)
	game.Play(2, 1)
	game.Play(0, 0)
	winner := game.Play(2, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_O_Win_Horizon_Third_Line_Should_Be_O_Winner(t *testing.T) {
	expected := "o Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(1, 0)
	game.Play(2, 1)
	game.Play(1, 1)
	game.Play(2, 0)
	game.Play(0, 1)
	winner := game.Play(2, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_O_Win_Vertical_First_Line_Should_Be_O_Winner(t *testing.T) {
	expected := "o Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(0, 1)
	game.Play(0, 0)
	game.Play(1, 1)
	game.Play(1, 0)
	game.Play(2, 1)
	winner := game.Play(2, 0)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Vertical_Second_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(1, 1)
	game.Play(0, 0)
	game.Play(2, 1)
	game.Play(1, 0)
	winner := game.Play(0, 1)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Vertical_Third_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(1, 2)
	game.Play(0, 0)
	game.Play(2, 2)
	game.Play(1, 0)
	winner := game.Play(0, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_O_Win_Right_Diagonal_Line_Should_Be_O_Winner(t *testing.T) {
	expected := "o Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(1, 0)
	game.Play(0, 0)
	game.Play(1, 2)
	game.Play(1, 1)
	game.Play(0, 2)
	winner := game.Play(2, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Left_Diagonal_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(2, 0)
	game.Play(0, 0)
	game.Play(1, 1)
	game.Play(2, 1)
	winner := game.Play(0, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_Mark_Full_Board_Shoud_Be_Tie(t *testing.T) {
	expected := "Tie"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(0, 2)
	game.Play(1, 1)
	game.Play(1, 0)
	game.Play(1, 2)
	game.Play(2, 1)
	game.Play(2, 0)
	winner := game.Play(2, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_Mark_Full_Board_And_X_Win_Horizon_Shoud_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := NewPlayer("mo", "x")
	player2 := NewPlayer("praw", "o")
	var player []Player
	player = append(player, player1, player2)
	game := NewGame(player, "x")

	game.Play(0, 0)
	game.Play(1, 0)
	game.Play(2, 1)
	game.Play(1, 1)
	game.Play(1, 2)
	game.Play(2, 2)
	game.Play(0, 1)
	game.Play(2, 0)
	winner := game.Play(0, 2)
	actual := winner

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
