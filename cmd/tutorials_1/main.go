package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Fibonacci function for demonstration
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	time.Sleep(10 * time.Millisecond)
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPU cores: %d\n", numCPU)

	numWorkers := numCPU // Set the number of worker goroutines to match the number of CPU cores
	totalFib := 100
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	numChannels := 20
	taskChs := make([]chan int, numChannels)
	for i := 0; i < numChannels; i++ {
		taskChs[i] = make(chan int)
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func(id int, taskCh chan int) {
			defer wg.Done()
			for {
				n, ok := <-taskCh
				if !ok {
					return
				}
				fmt.Printf("Worker %d on channel %d: Fibonacci(%d) = %d\n", id, id%numChannels, n, fibonacci(n))
			}
		}(i, taskChs[i])
	}

	// Distribute tasks evenly among channels
	for i := 0; i < totalFib; i++ {
		taskChs[i%numWorkers] <- i
	}

	// Close all channels to signal completion
	for _, ch := range taskChs {
		close(ch)
	}

	// Wait for all workers to complete
	wg.Wait()
}
