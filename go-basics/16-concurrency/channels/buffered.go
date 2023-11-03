package main

import (
	"fmt"
	"time"
)

// DemoBuffered demonstrates the buffered channels in Go
// Buffered channel is asynchronous to send multiple messages up to the specified size.
// After that it is blocked until someone reads the message.
func DemoBuffered() {

	// ----------------------------------------------------------
	// Create a Bufferred channel and can hold upto 3 messages
	// ----------------------------------------------------------
	msgQueue := make(chan string, 2)
	// Another channel to indicate routine completion
	done := make(chan bool) // "chan struct{}"" performs better than "chan bool"
	go func() {
		// It can accept 2 messages without any receiver
		msgQueue <- "msg1"
		fmt.Println("Sent msg1")
		msgQueue <- "msg2"
		fmt.Println("Sent msg2")

		// Third message is blocked until any message is read from the queue
		fmt.Println("Sending msg3 - expected to be blocked until read")
		msgQueue <- "msg3"
		fmt.Println("Sent msg3")

		// Indicate the completion
		done <- true
	}() // An anonymous function can also be a goroutine. Note the () at the end

	// Receive the message
	time.Sleep(5 * time.Second)
	fmt.Println("Reading messages")
	msg1 := <-msgQueue
	msg2 := <-msgQueue
	msg3 := <-msgQueue
	fmt.Println(msg1, msg2, msg3)

	// Wait for routine completion
	<-done
}


