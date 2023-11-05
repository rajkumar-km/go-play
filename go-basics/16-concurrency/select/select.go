/*
select demonstrates the use of select in Go

  - select multiplex multiple communication channels and wait for any send/receive communication.
  - It can also be used for non blocking send/receive by adding a default case.
  - Case ordering in select does not matter. It uses pseudo random that can distribute amoung
    multiple cases evenly. If you want some priority cases, then use separate select statements
    in outer blocks.
  - A nil channel is sometimes useful. Send/receive operations on a nil channel blocks forever.
    Using the nil channel in select is equivalent to disabling the case. It would never
    hit. So, we can conditionally disable some cases by setting nil value to the channel.
  - Empty select{} blocks forever
*/
package main

import (
	"fmt"
	"time"
)

func main() {
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

	// select can multiplex both send and receive operations
	ch := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case x := <-ch:
			fmt.Println("Receive: ", x)
		case ch <- i:
			fmt.Println("Send: ", i)
		}
	}

	// Note: Empty select{} blocks forever
}
