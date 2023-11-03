/*
channels demonstrates the use of Go channels
- Channels are the way to communicate between goroutines
- It can be considered like pipelines connecting multiple goroutines
- Multiple goroutines can access the same channel at the same time and it is concurrent.

Channel creations:

	ch := make(chan int) // unbuffered/syncronous channel
	ch := make(chan string, size) // buffered channel for specific size

Channel operations: Read/Write

	ch<- "hello"      // Send to channel
	v := <-ch         // Receive from channel
	v, ok := <-ch     // Receive from channel and the "ok" indicates if the channel is not closed
	for v := range ch // Range over the channel automatically detects the channel closure and ends

	Note that the operator "<-" is used for both read and write, but the channel variable is
	specified before the operator for write, and after for read.

Channel operations: Close

		close(ch)      // Close the channel

	- The close() operation sets the flag not to allow any further send operations.
	- Sending on to a closed channel causes panic.
	- Receive operations however works until there are values left on the channel and
	  finally returns the zero value of the type.
	- Closing a channel that is already closed causes panic.
	- Go has no mechanism to check if a channel is closed without reading a data from it.
	  So, in a multiple sender scenarios, who would close the channel?
	- We don't have to close the channel in all the scenarios. It is required only when we need to
	  inform the receivers that the data transfer is complete.
	- 1 sender: Sender can close the channel, receivers can stop reading from it
	- N senders: Don't have to close the channel in some cases. sync.Once can be preferred to
	  close the channel only once. A separate signalling channel can be used to indicate all
	  the senders that the channel is closed.
*/
package main

import "fmt"

func main() {
	// Simple unbuffered channel
	fmt.Println("--- DemoUnbuffered -------")
	DemoUnbuffered()

	// Buffered channel
	fmt.Println("\n--- DemoBuffered -------")
	DemoBuffered()

	// Channel directions
	fmt.Println("\n--- DemoChannelDirections -------")
	DemoChannelDirections()

	// Channel as pipelines
	fmt.Println("--- DemoChannelPipes -------")
	DemoChannelPipes()
}