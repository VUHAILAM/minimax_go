package source

import "fmt"

type Handler struct {
	board    *Board
	computer *Computer
	player   *User
}

func Init() *Handler {
	board := NewBoard()
	player := &User{}
	computer := &Computer{}
	fmt.Println("Wellcome to play Tic Tac Toe")
	board.Print()
	return &Handler{board, computer, player}
}

func (h *Handler) ChooseSymbol() {
	fmt.Println("Choose X or O")
	fmt.Println("1. X")
	fmt.Println("2. O")
	fmt.Println("Enter the number 1 or 2")
	var number int
	fmt.Scan(&number)
	switch number {
	case 1:
		h.player.symbol = PLAYER_MAX
		h.computer.symbol = PLAYER_MIN
	default:
		h.player.symbol = PLAYER_MIN
		h.computer.symbol = PLAYER_MAX
	}
	fmt.Println("Game starts, X go first!!")
}

func (h *Handler) HandleGame() {
	for turn := 0; turn < h.board.NumMove; turn++ {
		if turn%2 == 0 {
			if h.player.symbol == PLAYER_MAX {
				h.player.userMove(h.board)
				h.board.Print()
			} else {
				h.computer.computerMove(h.board)
				h.board.Print()
			}
		} else {
			if h.player.symbol == PLAYER_MAX {
				h.computer.computerMove(h.board)
				h.board.Print()
			} else {
				h.player.userMove(h.board)
				h.board.Print()
			}
		}
		if h.board.IsGameOver() {
			if h.board.IsWinner(PLAYER_MAX) {
				fmt.Println("X Win!!")
				return
			}

			if h.board.IsWinner(PLAYER_MIN) {
				fmt.Println("O Win!!")
				return
			}

			if h.board.NumMove == 0 {
				fmt.Println("Tie!!")
				return
			}
		}
	}
}
