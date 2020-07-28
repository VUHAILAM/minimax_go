package source

import (
	"fmt"
)

type User struct {
	symbol string
}

func (u *User) userMove(board *Board) {
	var x, y int
	fmt.Println("Enter the coordinates x and y")
	fmt.Scan(&x)
	fmt.Scan(&y)
	x--
	y--
	if board.Grid[x][y] == "*" {
		board.Grid[x][y] = u.symbol
	} else {
		fmt.Println("x and y were taken. Enter again!!")
		fmt.Scan(&x)
		fmt.Scan(&y)
	}
}
