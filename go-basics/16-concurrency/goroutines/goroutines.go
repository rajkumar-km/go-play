/*
goroutines demonstrates simple goroutine in Go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Main routine starts

	// spinner becomes a new goroutine
	go spinner(100 * time.Millisecond)

	// Main routine runs a job
	res := fib(45)
	fmt.Printf("\r45th fibonacci number is: %d\n", res)

	// Exit of main routine, also kills other goroutines
}

// fib returns the xth fibonacci number
func fib(x uint64) uint64 {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// spinner displays a loading text animation
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
