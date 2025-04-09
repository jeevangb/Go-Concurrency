package main

import (
	"fmt"
	"sync"
)

// This is a simple producer-consumer example in Go using channels and goroutines.
const numProducers = 2
const numConsumers = 3
const numJobs = 20

func main() {
	jobs := make(chan int, numJobs)
	var prodwg sync.WaitGroup
	var conswg sync.WaitGroup
	for i := 0; i < numProducers; i++ {
		prodwg.Add(1)
		go Producer(jobs, &prodwg)
	}
	for i := 0; i < numConsumers; i++ {
		conswg.Add(1)
		go Consumer(i, jobs, &conswg)
	}

	prodwg.Wait()
	close(jobs) // Close the jobs channel after all producers are done
	conswg.Wait()
}

func Producer(jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
}
func Consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Consumer", id, "processing job", job)
	}
}
