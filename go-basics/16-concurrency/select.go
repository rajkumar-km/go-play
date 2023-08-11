package main

import (
	"fmt"
	"time"
)

// DemoSelect demonstrates the use of select in Go
func DemoSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine that sends a message on channel c1
	go func() {
		c1 <- "msg1"
	}()

	// Select is blocked and waiting for a message from any one channel
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case msg1 := <-c1:
		fmt.Println(msg1)
	}

	// Non blocking send via select/default
	// Just add a default case in select to make it non blocking
	select {
	case c1 <- "msg1":
		fmt.Println("Sent message")
	default:
		fmt.Println("No message is sent")
	}

	// Non blocking recv via select/default
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	default:
		fmt.Println("No message is received")
	}

	// Using timeout in select
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out waiting for the message")
	}

	time.Sleep(1 * time.Second) // Sleep 1 second
}
