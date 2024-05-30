package sqr

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func highProcessingAlgorithm(workerID int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// Perform high processing algorithm
		result := j * j
		time.Sleep(100 * time.Millisecond) // Simulate processing time

		// Send result back to results channel
		results <- result
		fmt.Printf("Worker %d processed job %d\n", workerID, j)
	}
}

func main() {
	numWorkers := runtime.NumCPU() // Get number of CPU cores
	//fmt.Printf("Number of CPU cores: %d\n", numWorkers)

	jobs := make(chan int, 12)
	results := make(chan int, 1)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go highProcessingAlgorithm(w, jobs, results, &wg)
	}

	// Send jobs to workers
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to indicate no more jobs

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel after all workers are done
	}()

	// Collect results from workers
	for res := range results {
		fmt.Println("Result:", res)
	}
}
