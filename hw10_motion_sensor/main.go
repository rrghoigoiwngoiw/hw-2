package sensor

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Sensor() {
	sensorData := make(chan int)
	processedData := make(chan float64)

	go func() {
		defer close(sensorData)
		start := time.Now()
		for time.Since(start) < 1*time.Minute {
			n, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				panic(err) // Обработка ошибки
			}
			data := int(n.Int64()) // Преобразование в int
			sensorData <- data
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer close(processedData)
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
