/*
fetch_url fetches the given URL and prints the output

Go standard library has all the network related packages grouped under "net".
So, the "net/http" package is useful to establish http connections
  - http.Get(url) returns the response and error if any
  - The response.Body contains the stream to read the output
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Read the command line arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: fetch_url <URL>\n")
		os.Exit(1)
	}
	url := os.Args[1]

	// Send the GET request to given URL and receive the response
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get error: %v\n", err)
		os.Exit(1)
	}

	// Read the response and display in stdout
	n, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	// Close the response body to avoid leaking resources
	err = resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "close error: %v\n", err)
	}

	fmt.Printf("(%d bytes)", n)
}
