package main

import (
	"fmt"

	"github.com/nomannaq/Gambit-GolangChessBot/game"
)

func main() {
	fmt.Println("Welcome to Golang Chess Bot!")
	fmt.Println("You are playing as White. Enter moves in algebraic notation (e.g., e2e4).")
	game.PlayGame()
}
