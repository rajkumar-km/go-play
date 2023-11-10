/*
imit explain the pattern for limiting Go concurrency

Creating too much of gorutines sometimes causes resource constraints. For instance,
while doing the file operations, it causes "too many open files" error. This can be
solved by limiting the number of concurrent goroutines. This can be simply done by
having a buffered channel of desired size that issues a token.
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
	err := MakeThumbnails3(filenames)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MakeThumbnails3 receives the filenames in a channel and generates thumbnail images
// Only limited number of goroutines are active here to prevent from "too many open files" error
func MakeThumbnails3(filenames <-chan string) error {
	token := make(chan struct{}, 3) // Keep only 3 goroutines at a time
	errors := make(chan error)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1) // increment the counter
		go func(file string) {
			// Tokens are issues to only limited routines (using a buffered channel)
			// Other goroutines waits here until some routine release the token
			token <- struct{}{}        // acquire token
			defer func() { <-token }() // release token

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
