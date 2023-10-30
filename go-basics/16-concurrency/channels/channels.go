/*
channels demonstrates the use of Go channels
*/
package main

import (
	"fmt"
	"time"
)

// echo prints a message three times with 10ms interval
// Notifies the caller by sending "done" to channel ch.
func echo(ch chan string, msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 10)
	}
	ch <- "done" // Notify caller
}

func main() {
	// ----------------------------------
	// 1. Create a channel of type string
	// Note: Channels are goroutine safe, so no Mutex required
	// ----------------------------------
	// The channel is unbufferred by default. It is also called synchronous channel.
	// Meaning the send is not complete until the receiver reads the message
	// The receive also blocked until someone sends the message to channel
	messages := make(chan string)

	// From a seperate goroutine, Post a message "ping" to the channel
	go echo(messages, "ping")
		
	// Receive the message from the channel posted by another goroutine
	// The following call blocks until another goroutine writes the msg
	msg := <-messages
	fmt.Println(msg)

	// ----------------------------------------------------------
	// 2. Create a Bufferred channel and can hold upto 3 messages
	// ----------------------------------------------------------
	msgQueue := make(chan string, 2)
	// Another bufferred channel to indicate routine completion
	done := make(chan bool, 1)
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

	// ---------------------------
	// 3. Channel directions
	// See sendMsg() and receiveMsg() and note the type of arguments used to pass the channel
	// ---------------------------
	sendMsg(msgQueue)
	receiveMsg(msgQueue)
}

// A channel argument can be specified only for sending messages
func sendMsg(mQ chan<- string) {
	mQ <- "ping1"
	mQ <- "ping2"
}

// A channel argument can be specified only for receiving messages
func receiveMsg(mQ <-chan string) {
	fmt.Println(<-mQ, <-mQ)
}
