package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}
