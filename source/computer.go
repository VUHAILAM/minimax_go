package source

type Computer struct {
	board  *Board
	symbol string
	score  int
}

func (c *Computer) computerMove() {
	x := 0
	y := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if c.board.Grid[i][j] == "*" {
				c.board.Grid[i][j] = c.symbol
				if c.symbol == PLAYER_MAX {
					scoreMiniMax := c.board.ImplementMinimax(0, PLAYER_MIN)
					if scoreMiniMax > c.score {
						c.score = scoreMiniMax
						x = i
						y = j
					}
				} else {
					scoreMiniMax := c.board.ImplementMinimax(0, PLAYER_MAX)
					if scoreMiniMax < c.score {
						c.score = scoreMiniMax
						x = i
						y = j
					}
				}
			}
		}
	}
	c.board.Grid[x][y] = c.symbol
}
