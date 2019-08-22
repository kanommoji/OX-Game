package ox_test

import (
	"ox"
	"testing"
)

func Test_Play_Input_X_Mark_Horizon_First_Line_Should_Be_X_Winner(t *testing.T) {
	expected := "x Winner"

	player1 := ox.NewPlayer("mo", "x")
	player2 := ox.NewPlayer("praw", "o")
	game := ox.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 1, 1)
	game.Play(player1, 0, 1)
	game.Play(player2, 1, 0)

	actual := game.Play(player1, 0, 2)

	if expected != actual {
		t.Errorf("Expect %s bot got %s", expected, actual)
	}

}
