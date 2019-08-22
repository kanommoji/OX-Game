package ox_test

import (
	"ox"
	"testing"
)

func Test_X_Win_By_Horizon_First_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := ox.NewPlayer("mo", "x")
	player2 := ox.NewPlayer("praw", "o")
	game := ox.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 1, 1)
	game.Play(player1, 0, 1)
	game.Play(player2, 1, 0)
	winner := game.Play(player1, 0, 2)

	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_X_Win_Horizon_Second_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"
	player1 := ox.NewPlayer("mo", "x")
	player2 := ox.NewPlayer("praw", "o")
	game := ox.NewGame(player1, player2, "x")

	game.Play(player1, 1, 0)
	game.Play(player2, 0, 1)
	game.Play(player1, 1, 1)
	game.Play(player2, 0, 0)
	winner := game.Play(player1, 1, 2)

	actual := winner

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
