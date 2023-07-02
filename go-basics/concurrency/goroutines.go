package concurrency

import (
	"fmt"
	"time"
)

var ch chan struct{} // struct{} is smaller than bool

// printThrice prints a message three times with 10ms interval
func printThrice(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 10)
	}
	ch <- struct{}{} // Notify caller
}

// PlayGoroutines demonstrates the goroutines
func PlayGoroutines() {
	// Goroutine creates a new routine for concurrency
	ch = make(chan struct{})
	go printThrice("goroutine")

	<-ch // Wait
}
