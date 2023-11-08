package main

import "fmt"

// DemoChannelDirections demonstrates directions of channels in function arguments
// This can be specified in function arguments to restrict only for send or receive operations
// Example:
//	func sendMsg(mQ chan<- string) // send only
//	func receiveMsg(mQ <-chan string) // receive only
func DemoChannelDirections() {
	msgQueue := make(chan string, 2)

	// Directions are useful when:
	// - passing channel as arguments
	// - storing channels in a data structure that is used either for sending or receiving

	// See sendMsg() and receiveMsg() and note the type of arguments used to pass the channel
	sendMsg(msgQueue)
	receiveMsg(msgQueue)

	// Here, a map structure can store a channel as send only for broadcasting
	type client chan<- string
	broadcasterClients := make(map[client]bool)
	broadcasterClients[msgQueue] = true

	// So, a channel can be intermediately used only for sending or receiving for safety,
	// and indicating the intent of how it should be used.
	// However, creating unidirectional channel does not help although it is allowed
	ch := make(chan<- int, 1)
	ch <- 1
	// x := <-ch // compile error: can not receiving from send only channel
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