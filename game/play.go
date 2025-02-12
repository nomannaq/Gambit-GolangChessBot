package game

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

// PlayGame starts a new chess game
func PlayGame() {
	game := chess.NewGame()

	for game.Outcome() == chess.NoOutcome {
		fmt.Println(game.Position().Board().Draw())

		if game.Position().Turn() == chess.White {
			handlePlayerMove(game)
		} else {
			handleAIMove(game)
		}
	}

	fmt.Println(game.Position().Board().Draw())
	fmt.Println("Game Over! Result:", game.Outcome())
}

// Handle player's move
func handlePlayerMove(game *chess.Game) {
	fmt.Print("Enter your move (e.g., e2e4): ")
	var userMove string
	fmt.Scanln(&userMove)

	// Debugging: Print out the user input
	fmt.Println("User move:", userMove)

	// Check the valid moves for the current position
	validMoves := game.ValidMoves()
	fmt.Println("Valid moves:", validMoves)

	// Decode the move using AlgebraicNotation
	move, err := chess.AlgebraicNotation{}.Decode(game.Position(), userMove)
	if err != nil {
		// Debugging: Show the error
		fmt.Println("Error decoding move:", err)
		fmt.Println("Invalid move, try again.")
		handlePlayerMove(game)
		return
	}

	// Debugging: Show the decoded move
	fmt.Println("Decoded move:", move)

	// Make the move
	game.Move(move)
}

// Handle AI's move
func handleAIMove(game *chess.Game) {
	fmt.Println("Bot is thinking...")

	legalMoves := game.ValidMoves()
	if len(legalMoves) == 0 {
		log.Println("Bot has no legal moves! Stalemate or Checkmate.")
		return
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize RNG
	randomMove := legalMoves[rnd.Intn(len(legalMoves))]    // Pick a random move

	game.Move(randomMove)
	fmt.Println("Bot moved:", randomMove)
}
