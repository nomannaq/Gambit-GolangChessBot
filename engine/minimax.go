package engine

import "github.com/notnil/chess"

// MinimaxWithAlphaBeta applies minimax with alpha-beta pruning
func MinimaxWithAlphaBeta(game *chess.Game, depth int, alpha, beta int, maximizing bool) (int, *chess.Move) {
	if depth == 0 || game.Outcome() != chess.NoOutcome {
		return EvaluateBoard(game.Position().Board()), nil
	}

	var bestMove *chess.Move
	legalMoves := game.ValidMoves()

	if maximizing {
		bestScore := -99999
		for _, move := range legalMoves {
			newGame := game.Clone()
			newGame.Move(move)
			score, _ := MinimaxWithAlphaBeta(newGame, depth-1, alpha, beta, false)
			if score > bestScore {
				bestScore = score
				bestMove = move
			}
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break // Beta cut-off
			}
		}
		return bestScore, bestMove
	} else {
		bestScore := 99999
		for _, move := range legalMoves {
			newGame := game.Clone()
			newGame.Move(move)
			score, _ := MinimaxWithAlphaBeta(newGame, depth-1, alpha, beta, true)
			if score < bestScore {
				bestScore = score
				bestMove = move
			}
			beta = min(beta, bestScore)
			if beta <= alpha {
				break // Alpha cut-off
			}
		}
		return bestScore, bestMove
	}
}

// Helper functions for alpha-beta pruning
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
