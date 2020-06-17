package main

import (
	"fmt"
)

const (
	SIZE := 3
	PLAYER1 := "X"
	PLAYER2 := "O"
)

type Board struct {
	Grid[SIZE][SIZE] string
	NumMove int
}

func (b *Board) NewBoard() {
	b.Grid := [3][3]string{
		{"*", "*", "*"},
		{"*", "*", "*"},
		{"*", "*", "*"},
	}
	b.NumMove = SIZE*SIZE
}

func (b *Board) Print() {
	for i:= 1; i <= SIZE; i++ {
		fmt.Print(i+ " ")
	}
	for i:= 0; i < SIZE; i++ {
		fmt.Print((i+1)+" ")
		for j := 0; j < SIZE; j++ {
			fmt.Print(Grid[i][j]+ "|")
		}
		fmt.Println();
		fmt.Print("  ")
		for j := 0; j < SIZE; j++ {
			fmt.Print("- ")
		}
		fmt.Println()
	}
}

func (b Board) CheckVerticalWin(symbol string) bool {
	for i := 0; i < SIZE; i++ {
		column := [3]string{b.Grid[0][i], b.Grid[1][i], b.Grid[2][i]}
		if column == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b Board) CheckHorizontalWin(symbol string) bool {
	for i := 0; i < 3; i++ {
		row := b.Grid[i]
		if row == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b Board) CheckDiagonalWin(symbol string) bool {
	ddiag := [3]string{b.Grid[0][0], b.Grid[1][1], b.Grid[2][2]}
	udiag := [3]string{b.Grid[0][2], b.Grid[1][1], b.Grid[2][0]}
	winner := [3]string{symbol, symbol, symbol}
	if udiag == winner || ddiag == winner {
		return true
	}
	return false
}

func (b Board) IsWinner(symbol string) bool {
	// Checks if the game has been won by the passed symbol
	if b.verticalWin(symbol) || b.horizontalWin(symbol) || b.diagonalWin(symbol) {
		return true
	}
	return false
}

func (b Board) IsGameOver() bool {
	// Checks if the game is over (either player won or a draw)
	if b.Winner(PLAYER1) || b.Winner(PLAYER2) || b.TotalMoves == SIZE*SIZE {
		return true
	}
	return false
}
