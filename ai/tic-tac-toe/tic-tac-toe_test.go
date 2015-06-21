package tic_tac_toe

import (
	"testing"
)

func TestGame1NextMove1(t *testing.T) {
	board := []string{"___", "___", "___"}
	actual := nextMove("X", board)
	expected := "1 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame1NextMove2(t *testing.T) {
	board := []string{"O__", "_X_", "___"}
	actual := nextMove("X", board)
	expected := "1 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame1NextMove3(t *testing.T) {
	board := []string{"O__", "XXO", "___"}
	actual := nextMove("X", board)
	expected := "2 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame1NextMove4(t *testing.T) {
	board := []string{"OO_", "OXX", "X__"}
	actual := nextMove("X", board)
	expected := "0 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame1NextMove5(t *testing.T) {
	board := []string{"OOX", "OXX", "XO_"}
	actual := nextMove("X", board)
	expected := "2 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame2NextMove1(t *testing.T) {
	board := []string{"___", "___", "___"}
	actual := nextMove("X", board)
	expected := "1 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame2NextMove2(t *testing.T) {
	board := []string{"_O_", "_X_", "___"}
	actual := nextMove("X", board)
	expected := "1 0"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame2NextMove3(t *testing.T) {
	board := []string{"OO_", "XX_", "___"}
	actual := nextMove("X", board)
	expected := "1 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame3NextMove1(t *testing.T) {
	board := []string{"___", "___", "___"}
	actual := nextMove("X", board)
	expected := "1 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame3NextMove2(t *testing.T) {
	board := []string{"___", "OX_", "___"}
	actual := nextMove("X", board)
	expected := "0 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame3NextMove3(t *testing.T) {
	board := []string{"_X_", "OX_", "_O_"}
	actual := nextMove("X", board)
	expected := "0 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame3NextMove4(t *testing.T) {
	board := []string{"_XX", "OX_", "OO_"}
	actual := nextMove("X", board)
	expected := "0 0"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame4NextMove1(t *testing.T) {
	board := []string{"OXO", "OXX", "XO_"}
	actual := nextMove("X", board)
	expected := "2 2"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestGame4NextMove2(t *testing.T) {
	board := []string{"XXO", "OOX", "X__"}
	actual := nextMove("X", board)
	expected := "2 1"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasLastEmptyCell1(t *testing.T) {
	board := []string{"OXO", "OXX", "XO_"}
	actual := hasLastEmptyCell(board)
	expected := true
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
