/*
fetch_urls fetches the given URLs concurrently and prints the output
The total time of execution is equal to the longest time taken for a single URL
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Create a channel for goroutines communication
	var ch chan string = make(chan string)

	// Read the URLs from command line arguments
	for _, url := range os.Args[1:] {
		// Format the URL to include http:// prefix
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		// Fetch the URL in a separate routine
		go fetchUrl(url, ch)
	}

	// Finally, print the results from channel
	for i := 0; i < (len(os.Args)-1); i++ {
		fmt.Println(<-ch)
	}

	fetchDur := time.Since(start)
	fmt.Println("Total fetch time:", fetchDur)
}

// fetchUrl performs GET on the given URL and writes the error/response time
// in the given channel
func fetchUrl(url string, ch chan string) {
	start := time.Now()

	// Send the GET request to given URL and receive the response
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("http get error: %v", err)
		return
	}

	// Create a file and copy the response
	urlWithoutPrefix := strings.TrimPrefix(url, "http://")
	urlWithoutPrefix = strings.TrimPrefix(urlWithoutPrefix, "https://")
	outfilename := filepath.Join(os.TempDir(), urlWithoutPrefix)
	fmt.Println(outfilename)
	outfile, err := os.Create(outfilename)
	nbytes, err := io.Copy(outfile, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("read error: %v", err)
		return
	}

	// Close the response body to avoid leaking resources
	err = resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "close error: %v", err)
	}

	dur := time.Since(start)
	ch <- fmt.Sprintf("%.2fs %7d %s", dur.Seconds(), nbytes, url)
}