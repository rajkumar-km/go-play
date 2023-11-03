package main

import (
	"fmt"
	"time"
)

// DemoUnbuffered demonstrates unbuffered channels in Go
// It is also called synchronous channel since it blocks both send and recv operations
// when the other side is ready
func DemoUnbuffered() {
	// ----------------------------------
	// 1. Create a channel of type string
	// Note: Channels are goroutine safe, so no Mutex required
	// ----------------------------------
	// The channel is unbufferred by default. It is also called synchronous channel.
	// Meaning the send is not complete until the receiver reads the message
	// The receive also blocked until someone sends the message to channel
	ch := make(chan string)

	// From a seperate goroutine, Post a message "ping" to the channel
	go echo(ch, "ping")

	// Receive the message from the channel posted by another goroutine
	// The following call blocks until another goroutine writes the msg
	msg := <-ch
	fmt.Println("First msg", msg)

	// Receive message, but also check if the channel is closed
	msg, ok := <-ch
	if !ok {
		// ok will be false if the channel is closed
		// msg will be "" (the zero value for the string type)
		fmt.Println("channel closed")
		return
	}
	fmt.Println("Second msg", msg)

	// Use range and receive the messages as long as the channel is open
	for msg := range ch {
		fmt.Println("All other msgs", msg)
	}
}

// echo prints a message three times with 10ms interval
// Notifies the caller by sending "done" to channel ch.
func echo(ch chan string, msg string) {
	for i := 1; i <= 5; i++ {
		ch<- fmt.Sprintf("%d. %s", i, msg)
		time.Sleep(time.Millisecond * 10)
	}
	close(ch) // Notify caller
}