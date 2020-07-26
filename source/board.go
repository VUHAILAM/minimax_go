package source

import (
	"fmt"
	"math"
)

const (
	SIZE       = 3
	PLAYER_MAX = "X"
	PLAYER_MIN = "O"
)

type Board struct {
	Grid    [SIZE][SIZE]string
	NumMove int
}

func NewBoard() *Board {
	g := [3][3]string{
		{"*", "*", "*"},
		{"*", "*", "*"},
		{"*", "*", "*"},
	}
	num_move := SIZE * SIZE
	return &Board{g, num_move}
}

func (b *Board) Print() {
	for i := 1; i <= SIZE; i++ {
		fmt.Print(i, " ")
	}
	for i := 0; i < SIZE; i++ {
		fmt.Print((i + 1), " ")
		for j := 0; j < SIZE; j++ {
			fmt.Print(b.Grid[i][j], "|")
		}
		fmt.Println()
		fmt.Print("  ")
		for j := 0; j < SIZE; j++ {
			fmt.Print("- ")
		}
		fmt.Println()
	}
}

func (b *Board) CheckVerticalWin(symbol string) bool {
	for i := 0; i < SIZE; i++ {
		column := [3]string{b.Grid[0][i], b.Grid[1][i], b.Grid[2][i]}
		if column == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b *Board) CheckHorizontalWin(symbol string) bool {
	for i := 0; i < 3; i++ {
		row := b.Grid[i]
		if row == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b *Board) CheckDiagonalWin(symbol string) bool {
	ddiag := [3]string{b.Grid[0][0], b.Grid[1][1], b.Grid[2][2]}
	udiag := [3]string{b.Grid[0][2], b.Grid[1][1], b.Grid[2][0]}
	winner := [3]string{symbol, symbol, symbol}
	if udiag == winner || ddiag == winner {
		return true
	}
	return false
}

func (b *Board) IsWinner(symbol string) bool {
	// Checks if the game has been won by the passed symbol
	if b.CheckVerticalWin(symbol) || b.CheckHorizontalWin(symbol) || b.CheckDiagonalWin(symbol) {
		return true
	}
	return false
}

func (b *Board) IsGameOver() bool {
	// Checks if the game is over (either player won or a draw)
	if b.IsWinner(PLAYER_MAX) || b.IsWinner(PLAYER_MIN) || b.NumMove == 0 {
		return true
	}
	return false
}

func (b *Board) ImplementMinimax(depth int, player string) int {
	var mark int
	if b.IsWinner(PLAYER_MAX) {
		mark = 11 - depth
		return mark
	} else if b.IsWinner(PLAYER_MIN) {
		mark = depth - 11
		return mark
	}
	if b.IsGameOver() {
		return 0
	}
	if player == PLAYER_MAX {
		mark = -20
	} else {
		mark = 20
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.Grid[i][j] == "*" {
				if player == PLAYER_MAX {
					b.Grid[i][j] = PLAYER_MAX
					mark = int(math.Max(float64(mark), float64(b.ImplementMinimax(depth+1, PLAYER_MIN))))
				} else {
					b.Grid[i][j] = PLAYER_MIN
					mark = int(math.Min(float64(mark), float64(b.ImplementMinimax(depth+1, PLAYER_MAX))))
				}
				b.Grid[i][j] = "*"
			}
		}
	}
	return mark
}
