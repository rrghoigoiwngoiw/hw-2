package sensor

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Sensor() {
	processedData := make(chan float64)
	var sensorData []int
	go func() {
		start := time.Now()
		for time.Since(start) < 1*time.Minute {
			n, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				panic(err)
			}
			data := int(n.Int64())
			sensorData = append(sensorData, data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer close(processedData)
		count := 0
		sum := 0

		for data := range sensorData {
			sum += data
			count++
			if count == 10 {
				average := float64(sum) / 10.0
				processedData <- average
				count = 0
				sum = 0
			}
		}
	}()

	for result := range processedData {
		fmt.Printf("Среднее значение: %.2f\n", result)
	}

	fmt.Println("Завершение обработки данных.")
}
