package main

import (
	"fmt"
	"time"
)

// DemoChannelPipes demonstrates using channel as pipelines between goroutines
// Lets see an example communication between:
//  Routine 1 => (inputs channel) => Routine 2 => (squares channel) => Routine 3
func DemoChannelPipes() {
	done := make(chan struct{}) // signalling channel to exit DemoChannelPipes

	// Routine 1: Produces sequence of numbers in inputs channel
	inputs := make(chan int)
	go func() {
		for i := 0; i < 30; i++ {
			inputs <- i
			time.Sleep(10 * time.Millisecond)
		}
		close(inputs) // notify for completion
	}()

	// Routine 2: Read the inputs generated by routine 1 and squares them
	squares := make(chan int)
	go func() {
		for v := range inputs {
			squares <- (v*v)
		}
		close(squares) // notify for completion
	}()

	// Routine 3: Read the squares generated by routine 2 and print them
	go func() {
		for v := range squares {
			fmt.Println(v)
		}
		done <- struct{}{} // signal overall routine for completion
	}()

	<-done // Wait for Routine 3 completion
}