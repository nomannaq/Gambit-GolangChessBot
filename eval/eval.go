package eval

import "github.com/nomannaq/Gambit-GolangChessBot/board"

// Evaluate calculates the board's score
func Evaluate(b board.Board) int {
	score := 0
	pieceValues := map[int]int{
		board.Pawn: 100, board.Knight: 300, board.Bishop: 300,
		board.Rook: 500, board.Queen: 900, board.King: 10000,
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := b[y][x]
			if piece == board.Empty {
				continue
			}

			value := pieceValues[piece & ^board.ColorMask]
			if piece&board.ColorMask == board.White {
				score += value
			} else {
				score -= value
			}
		}
	}
	return score
}
