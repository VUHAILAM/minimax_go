package source

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	expectedResult := [3][3]string{
		{"*", "*", "*"},
		{"*", "*", "*"},
		{"*", "*", "*"},
	}
	result := b.Grid
	if result != expectedResult {
		t.Errorf("got %.s want %.s", result, expectedResult)
	}
}

func TestCheckVerticalWin(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 3; i++ {
		b.Grid[i][0] = "X"
	}
	result := b.CheckVerticalWin("X")
	expectedResult := true
	if result != expectedResult {
		t.Errorf("got %t want %t", result, expectedResult)
	}
}

func TestCheckHorizontalWin(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 3; i++ {
		b.Grid[0][i] = "X"
	}
	result := b.CheckHorizontalWin("X")
	expectedResult := true
	if result != expectedResult {
		t.Errorf("got %t want %t", result, expectedResult)
	}
}

func TestCheckDiagonalWin(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 3; i++ {
		b.Grid[i][i] = "X"
	}
	result := b.CheckDiagonalWin("X")
	expectedResult := true
	if result != expectedResult {
		t.Errorf("got %t want %t", result, expectedResult)
	}
}

func TestIsWinner(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 3; i++ {
		b.Grid[i][i] = "X"
	}
	result := b.IsWinner("X")
	expectedResult := true
	if result != expectedResult {
		t.Errorf("got %t want %t", result, expectedResult)
	}
}

func TestImplementMinimax(t *testing.T) {
	testCase := []struct {
		Board  Board
		Depth  int
		Player string
		Result int
	}{
		{
			Board: Board{
				Grid: [SIZE][SIZE]string{
					{"O", "O", "X"},
					{"X", "X", "O"},
					{"X", "O", "*"},
				},
				NumMove: SIZE*SIZE - 8,
			},
			Depth:  0,
			Player: PLAYER_MAX,
			Result: 11,
		},
	}
	for _, tc := range testCase {
		res := tc.Board.ImplementMinimax(tc.Depth, tc.Player)
		if res != tc.Result {
			t.Errorf("got %d want %d", res, tc.Result)
		}
	}
}
