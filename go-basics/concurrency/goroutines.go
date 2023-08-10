package main

import (
	"fmt"
	"time"
)

// DemoGoroutines demonstrates the goroutines
func DemoGoroutines() {
	var ch chan struct{} // struct{} is smaller than bool

	// Goroutine creates a new routine for concurrency
	ch = make(chan struct{})
	go printThrice(ch, "goroutine")

	<-ch // Wait
}

// printThrice prints a message three times with 10ms interval
func printThrice(ch chan struct{}, msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 10)
	}
	ch <- struct{}{} // Notify caller
}

