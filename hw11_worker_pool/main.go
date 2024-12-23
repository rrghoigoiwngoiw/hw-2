package pool

import (
	"fmt"
	"sync"
)

func Pool() int {
	var (
		wg      sync.WaitGroup
		counter struct {
			mu    sync.Mutex
			count int
		}
	)

	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			counter.mu.Lock()
			counter.count++
			counter.mu.Unlock()

			fmt.Printf("Worker %d done\n", id)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return counter.count
}
