package chessboard

import (
	"errors"
	"strings"
)

// переделал функцию чтобы ее можно было легко тестировать
func GenerateChessboard(size int) (string, error) {
	if size <= 0 {
		return "", errors.New("размер доски должен быть положительным")
	}

	var builder strings.Builder
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if (y+x)%2 == 0 {
				builder.WriteString("# ")
			} else {
				builder.WriteString("  ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String(), nil
}
