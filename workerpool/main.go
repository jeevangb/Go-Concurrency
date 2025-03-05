package main

import (
	"fmt"
	"sync"
)

func main() {
	const numWorkers = 3
	const numJobs = 60

	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, result, &wg)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(result)
	}()
	for res := range result {
		fmt.Println("Result :", res)
	}

}
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}
