/*
buffered demonstrates the concurrency pattern to loop in parallel using a buffered channel

Consider a operation to make thumbnail image of several files. This can be concurrent
using goroutines as follows. We want to return any error as soon as encounted. We tend to
use an unbuffered channel for communication, but it would cause gorutines to block once
the function returns on first error without receiving all messages. This is called
goroutine leak. There are two ways we can address this issue:
 1. Use buffered channel of sufficient size (if the length can be determined)
 2. Use a separate closer goroutine to wait and close other routines
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s file-1 [file-2] [file-3] [file-n] ", os.Args[0])
		os.Exit(1)
	}

	// Run the jobs in main routine and exit on first error
	err := MakeThumbnails(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MakeThumbnails generate thumbnail images for the given files
// Note that we create a buffered channel in this case. Choosing an unbuffered channel would
// cause goroutine leak in case if we return on any error.
func MakeThumbnails(filenames []string) error {
	errors := make(chan error, len(filenames))

	for _, f := range filenames {
		go func(file string) {
			// Integrate the thumbnal image functionality here
			// _, err := thumbnail.ImageFile(file)
			time.Sleep(3 * time.Second)
			thumbfile := strings.TrimSuffix(file, filepath.Ext(file))
			fmt.Printf("Created thumbnail: %s\n", thumbfile+".thumb.png")
			var err error = nil

			// Sending never blocks since we have sufficient buffer for each filename
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			// Free to return without receiving all the errors
			// Buffered channel does not block sending goroutines
			// However, the thumbnail routine runs in the background as long as the main
			// goroutine is alive
			return err
		}
	}

	return nil
}
