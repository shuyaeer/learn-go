package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	results <- id * 2
}

func main() {
	const numWorkers = 5
	var wg sync.WaitGroup
	results := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
