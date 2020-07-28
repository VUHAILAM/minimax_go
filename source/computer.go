package source

type Computer struct {
	symbol string
	score  int
}

func (c *Computer) computerMove(board *Board) {
	x := 0
	y := 0
	if c.symbol == PLAYER_MAX {
		c.score = -20
	} else {
		c.score = 20
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board.Grid[i][j] == "*" {
				board.Grid[i][j] = c.symbol
				if c.symbol == PLAYER_MAX {
					scoreMiniMax := board.ImplementMinimax(0, PLAYER_MIN)
					if scoreMiniMax > c.score {
						c.score = scoreMiniMax
						x = i
						y = j
					}
				} else {
					scoreMiniMax := board.ImplementMinimax(0, PLAYER_MAX)
					if scoreMiniMax < c.score {
						c.score = scoreMiniMax
						x = i
						y = j
					}
				}
			}
		}
	}
	board.Grid[x][y] = c.symbol
}
