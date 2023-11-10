/*
closer demonstrates the Go concurrency pattern to launch a closer Go routine that
waits for all other routines to avoid goroutines leak.

A buffered channel can be used instead but we need to know the maximum goroutnes
to set the buffer size. The following example receives the input through a channel,
so we can not determine the buffer size. See patterns/buffered for more details.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s file-1 [file-2] [file-3] [file-n] ", os.Args[0])
		os.Exit(1)
	}

	// Input channel to send the filenames
	filenames := make(chan string)
	go func() {
		for _, f := range os.Args[1:] {
			filenames <- f
		}
		close(filenames)
	}()

	// Run the jobs in main routine and exit on first error
	err := MakeThumbnails2(filenames)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MakeThumbnails2 receives the filenames in a channel and generates thumbnail images
// We use waitgroup to count the number of routines and wait for it. A separate closer
// goroutine is running while the function returns on first error.
func MakeThumbnails2(filenames <-chan string) error {
	errors := make(chan error)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1) // increment the counter
		go func(file string) {
			defer wg.Done() // decrement the counter once the routine is complete

			// Integrate the thumbnal image functionality here
			// _, err := thumbnail.ImageFile(file)
			time.Sleep(3 * time.Second)
			thumbfile := strings.TrimSuffix(file, filepath.Ext(file))
			fmt.Printf("Created thumbnail: %s\n", thumbfile+".thumb.png")
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
