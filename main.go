package main

import (
	src "minimax_go/source"
)

func main() {
	handler := src.Init()
	handler.ChooseSymbol()
	handler.HandleGame()
}
