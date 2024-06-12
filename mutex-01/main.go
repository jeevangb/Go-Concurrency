package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	var lock sync.Mutex

	//increment function
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	//decrement function
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	var arithemetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithemetic.Add(1)
		go func() {
			defer arithemetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithemetic.Add(1)
		go func() {
			defer arithemetic.Done()
			decrement()
		}()
	}
	arithemetic.Wait()
	fmt.Println("Arithemetic completed")
}
