package main

import "fmt"

func main() {
	// Unbuffered channel
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	// Buffered channel
	c := make(chan int, 2)

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}
