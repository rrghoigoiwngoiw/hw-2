package sensor

import (
	"testing"
	"time"
)

func runSensorTest(t *testing.T, randomValues []int, expectedAvg []float64) {
	t.Helper()

	sensorData := make(chan int, len(randomValues))
	processedData := make(chan float64)

	go func() {
		defer close(sensorData)
		for _, value := range randomValues {
			sensorData <- value
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		defer close(processedData)
		buffer := make([]int, 0, 10)
		for data := range sensorData {
			buffer = append(buffer, data)
			if len(buffer) == 10 {
				sum := 0
				for _, v := range buffer {
					sum += v
				}
				processedData <- float64(sum) / 10.0
				buffer = buffer[:0]
			}
		}
	}()

	results := make([]float64, 0, len(expectedAvg))
	for avg := range processedData {
		results = append(results, avg)
	}

	if len(results) != len(expectedAvg) {
		t.Fatalf("Expected %d averages, got %d", len(expectedAvg), len(results))
	}
	for i, expected := range expectedAvg {
		if results[i] != expected {
			t.Errorf("At index %d: expected %.2f, got %.2f", i, expected, results[i])
		}
	}
}

func TestSimpleAverage(t *testing.T) {
	randomValues := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	expectedAvg := []float64{55.0}
	runSensorTest(t, randomValues, expectedAvg)
}

func TestMultipleAverages(t *testing.T) {
	randomValues := []int{5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	expectedAvg := []float64{50.0, 55.0}
	runSensorTest(t, randomValues, expectedAvg)
}

func TestIncompleteBuffer(t *testing.T) {
	randomValues := []int{1, 2, 3, 4, 5}
	expectedAvg := []float64{}
	runSensorTest(t, randomValues, expectedAvg)
}

func TestZeroValues(t *testing.T) {
	randomValues := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedAvg := []float64{0.0}
	runSensorTest(t, randomValues, expectedAvg)
}
