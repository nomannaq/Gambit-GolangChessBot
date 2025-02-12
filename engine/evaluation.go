package engine

import "github.com/notnil/chess"

// Piece values for evaluation
var pieceValues = map[chess.PieceType]int{
	chess.Pawn:   1,
	chess.Knight: 3,
	chess.Bishop: 3,
	chess.Rook:   5,
	chess.Queen:  9,
	chess.King:   1000,
}

// EvaluateBoard calculates a basic material score for a given board position
func EvaluateBoard(board *chess.Board) int {
	score := 0
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := board.Piece(sq)
		if piece == chess.NoPiece {
			continue
		}
		value := pieceValues[piece.Type()]
		if piece.Color() == chess.White {
			score += value
		} else {
			score -= value
		}
	}
	return score
}
