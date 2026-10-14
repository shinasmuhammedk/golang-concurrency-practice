package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("worker %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

    wg.Add(3)
	for i := 1; i < 4; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}
