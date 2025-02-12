package move

import "fmt"

type Move struct {
	FromX, FromY int
	ToX, ToY     int
	Promotion    int
}

// ParseMove converts algebraic notation (e.g., "e2e4") to a Move
func ParseMove(moveStr string) (Move, error) {
	if len(moveStr) != 4 {
		return Move{}, fmt.Errorf("invalid move format")
	}
	fromX := int(moveStr[0] - 'a')
	fromY := 7 - int(moveStr[1]-'1')
	toX := int(moveStr[2] - 'a')
	toY := 7 - int(moveStr[3]-'1')
	return Move{FromX: fromX, FromY: fromY, ToX: toX, ToY: toY}, nil
}
