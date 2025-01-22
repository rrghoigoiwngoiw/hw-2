package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileFlag := flag.String("file", "", "файл откуда будем читать")
	levelFlag := flag.String("level", "info", "указывает уровень логов для анализа")
	outputFlag := flag.String("output", "", "файл в который будет записана статистика ")

	flag.Parse()

	// Проверяем переменные окружения, если флаг не задан
	file := *fileFlag
	if file == "" {
		file = os.Getenv("LOG_ANALYZER_FILE")
	}
	if file == "" {
		fmt.Println("Error: log file path is not specified. Use -file or set LOG_ANALYZER_FILE.")
		return
	}

	level := *levelFlag
	if level == "" {
		level = os.Getenv("LOG_ANALYZER_LEVEL")
		if level == "" {
			level = "info" // Значение по умолчанию
		}
	}

	output := *outputFlag
	if output == "" {
		output = os.Getenv("LOG_ANALYZER_OUTPUT")
		// Если output остается пустым, данные выводятся в stdout
	}

}
