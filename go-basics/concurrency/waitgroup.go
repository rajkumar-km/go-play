package main

import (
	"fmt"
	"sync"
)

// DemoWaitgroup demonstrates the waitgroup
func DemoWaitGroup() {
	// A wait group can be used to wait for goroutine completion
	var wg sync.WaitGroup

	// Goroutine with anonymous function
	wg.Add(1) // Adds wait group delta by 1
	go func(msg string) {
		defer wg.Done() // Decrements wait group by 1
		fmt.Println("anonymous goroutine")
	}("via anonymous goroutine")

	// Wait for routines completion
	// Note: This waits for only anonymous goroutine since we did not
	// use waitgroup for previous goroutine
	wg.Wait() // Blocks until the wait group delta is zero
}
