package main

import (
	"fmt"
	"sync"
)

func worker(jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result <- job * job
	}
}

func main() {
	jobs := make(chan int)
	result := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go worker(jobs, result, &wg)
	}

	go func() {
		for j := 1; j <= 6; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

    for value := range result{
        fmt.Println(value)
    }

}
