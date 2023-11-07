/*
cancellation demonstrates the common Go concurrency pattern to cancel multiple goroutines

Go does not support closing other goroutines from the parent routines, because this can leave
the goroutines states as it is. Instead, the goutine itself can implement some logic to exit.

A single goroutine may be cancelled by having a done channel. Sending a value to the done
channel can be the signal and the goroutine can exit as soon as it receives the message.

But what if there are hundreds of thousands of goroutines and want to cancel all by sending a
signal? The value on the channel is lost once the first goroutine reads. So, close the channel
instead of sending a value. All the goroutines receives a zero value when the channel is closed.
*/
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var cancel = make(chan struct{}) // Signal for cancelling all the jobs

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s file-1 [file-2] [file-3] [file-n] ", os.Args[0])
		os.Exit(1)
	}

	// Input channel to send the filenames
	filenames := make(chan string)
	go func() {
		for i := range os.Args[1:] {
			filenames <- os.Args[i]
		}
		close(filenames)
	}()

	// Signal cancel when the user hit ENTER
	go func() {
		fmt.Print("Press ENTER to abort")
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
		fmt.Println("Sent signal to abort all goroutines")
	}()

	// Run the jobs in main routine and exit on first error
	err := MakeThumbnails(filenames)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MakeThumbnails receives the filenames in a channel and generates thumbnail images
// Cancellation is supported by closing a "done" channel which broadcast to all routines
func MakeThumbnails(filenames <-chan string) error {
    token := make(chan struct{}, 2) // To keep only 2 goroutines at a time
    errors := make(chan error)
    var wg sync.WaitGroup

    for f := range filenames {
        wg.Add(1) // increment the counter
        go func(file string) {
            defer wg.Done() // decrement the counter once the routine is complete

            select {
            case token <- struct{}{}: // acquire token
                // Tokens are issues to only limited routines (using a buffered channel)
                // Other goroutines waits here until some routine release the token
			case <-cancel:
                // Received cancel signal, so do not spawn new jobs
				fmt.Println("aborted while waiting for token")
                return
            }            
            defer func() { <-token }() // release token
            
            // Depends on the operations, we may need to add this check in multiple places
            // to handle the cancellation signals
            if cancelled() {
				fmt.Println("aborted before thumbnail creation")
                return
            }

            // Integrate the thumbnal image functionality here
			// _, err := thumbnail.ImageFile(file)
			time.Sleep(3*time.Second)
			var err error = nil

            errors <- err
        }(f)
    }

    // closer routine wait for all other routines
    go func() {
        wg.Wait() // wait for all routines. Counter becomes zero
        close(errors)
    }()

    // Main routine return any errors as soon as it encounters
    for err := range errors {
        if err != nil {
            return err
        }
    }

    return nil
}

// cancelled returns true if the operations are cancelled, false otherwise.
func cancelled() bool {
    select {
        case <-cancel:
            return true
        default:
            return false
    }
}