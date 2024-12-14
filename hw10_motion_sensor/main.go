package sensor

import (
	"fmt"
	"math/rand"
	"time"
)

func Sensor() {
	sensorData := make(chan int)
	processedData := make(chan float64)

	go func() {
		defer close(sensorData) // Закрываем канал после завершения работы
		start := time.Now()
		for time.Since(start) < 1*time.Minute {
			data := rand.Intn(100)
			sensorData <- data
			time.Sleep(1 * time.Second)
		}
	}()

	// Горутина для обработки данных
	go func() {
		defer close(processedData) // Закрываем канал после завершения работы
		buffer := make([]int, 0, 10)
		for data := range sensorData {
			buffer = append(buffer, data)
			if len(buffer) == 10 {
				sum := 0
				for _, value := range buffer {
					sum += value
				}
				average := float64(sum) / 10.0
				processedData <- average
				buffer = buffer[:0]
			}
		}
	}()

	for result := range processedData {
		fmt.Printf("Среднее значение: %.2f\n", result)
	}

	fmt.Println("Завершение обработки данных.")
}
