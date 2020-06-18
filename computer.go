package main

import "math"

func implementMinimax(board *Board, depth int, player string) int {
	var mark int
	if board.IsWinner(PLAYER1) {
		mark = 11 - depth
		return mark
	} else if board.IsWinner(PLAYER2) {
		mark = depth - 11
		return mark
	}
	if board.IsGameOver() {
		return 0
	}
	if player == PLAYER1 {
		mark = -20
	} else {
		mark = 20
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == "*" {
				if player == PLAYER1 {
					board[i][j] = PLAYER1
					mark = math.Max(mark, implementMinimax(&board, depth+1, PLAYER2))
				} else {
					board = PLAYER2
					mark = math.Min(mark, implementMinimax(&board, depth+1, PLAYER1))
				}
				board[i][j] = "*"
			}
		}
	}
	return mark
}
