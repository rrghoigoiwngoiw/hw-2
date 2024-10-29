package main

import "fmt"

func main() {
	var chessSize int

	fmt.Print("Введите размер доски: ")
	_, err := fmt.Scan(&chessSize)
	if err != nil || chessSize <= 0 {
		fmt.Println("Ошибка: введите положительное целое число")
		return
	}

	for y := 0; y < chessSize; y++ {
		for x := 0; x < chessSize; x++ {
			if (y+x)%2 == 0 {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}
