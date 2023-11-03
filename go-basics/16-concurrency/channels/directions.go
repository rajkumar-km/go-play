package main

import "fmt"

// DemoChannelDirections demonstrates directions of channels in function arguments
// This can be specified in function arguments to restrict only for send or receive operations
// Example:
//	func sendMsg(mQ chan<- string) // send only
//	func receiveMsg(mQ <-chan string) // receive only
func DemoChannelDirections() {
	msgQueue := make(chan string, 2)

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