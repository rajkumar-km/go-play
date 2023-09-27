package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

// DemoErrors demonstrates error handling in Go
//   - Some functions always succeeds so they don't need to return errors.
//   - Some functions always returns single failure, so a bool can be used.
//   - An "error" type of argument is required if a function can return multiple errors.
//   - The other return values are undefined when an error is returned. However, some
//     functions may need to return partial results along with the error. Say a Read()
//     returning partial data with IO errors. Such case should be documented well.
//
// Go took the design decision to go with this error type instead of exceptions:
//   - Exceptions returns incomprehensible stacktrace, full of information about the
//     structure of the program, but lacking intelligable context about what went wrong.
//   - By contrast, Go programs use ordinary error handling with return values and
//     if statements. The demands more attention to be paid to error handling. This is
//     the primary reason.
//   - Go also has panic() and recover() which is a kind of exception handling, but
//     that is used to report truly unexpected exceptions that indicates a bug.
//
// The error type is an interface:
//   - A nil value is default and indicates no error.
//   - it has the method "Error() string" to get the error message.
//   - Package "errors" contains several helpers
//   - fmt.Errorf is handy to produce error type that uses Sprintf internally
//
// Error handling techniques:
//  1. Propagate the error
//     a. Return as it is
//     b. Add context to the error and form the chain
//  2. Retry for sometime
//  3. Exit with error code (main workflow)
//  4. Ignore the error continue with limited functionality
//  5. Ignore the error and proceed
//
// In general, get into the habit of handling errors returned by functions.
//   - If you deliberately ignore some errors, document it.
//   - Handle the error first before performing success action
//   - If return statement is used for error, then write the success action in the outer block.
//     Do not use else block and minimize in indent.
//   - So a function tends to handle multiple errors on top and return. Finally the sustance
//     of the function.
func DemoErrors() {
	// 1. Propagate the error
	fmt.Println("1. Propagate the error")
	links, err := fetchlinks("https://golang.org")
	if err != nil {
		// Ignore the error and proceed without printing links
		log.Printf("fetchlinks failed: %v", err)
	} else {
		log.Printf("Found %d links\n", len(links))
	}

	// 2. Retry
	fmt.Println("2. Retry on error")
	resp, err := getTimeout("https://golang.org", 5 * time.Second)
	if err != nil {
		// 3. Exit on error
		fmt.Println("3. Exit on error")
		log.Fatalf("Connection error: %v\n", err)
		// or use log.Printf with os.Exit(1)
	} else {
		data, err := io.ReadAll(resp)
		if err != nil {
			// 4. Ignore the error with reduced functionality (whatever is in the else case)
			fmt.Println("4. Ignore the error with reduced functionality")
			log.Printf("io.ReadAll failed: %v", err)
		} else {
			log.Printf("Read %d bytes\n", len(data))
		}
	}

	// 5. Ignore the error altogether and continue
	_, err = fmt.Fprintf(os.Stdout, "5. Ignore the error altogether and proceed\n")
	fmt.Println("Bye")
}

// get fetches a url and returns the response stream
// propagates if any error is returned by http.Get
func fetchlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// Propagate the error
		// - if the error is meaningful and provides context to the caller
		// get failed: Get "https://golang.orgg": dial tcp: lookup golang.orgg: no such host
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		// Add more context to the error
		// - If the error lacks some information, add them to form the chain of errors. Avoid
		//   capitalizing the message, avoid newlines, and be consistent with the message.
		// - For example: os.Open(f) adds the filename along with error message:
		//   "open example.txt: permission denied"
		// - The function is responsible to the error and the arguments related to it. Caller is
		//   is responsible the information that is not passed to the function.
		return nil, fmt.Errorf("parsing %s failed: %v", url, err)
	}

	links := visit(nil, doc)

	return links, nil
}

// visit finds the links in HTML document
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	if node.FirstChild != nil {
		// A recursive call to first child
		links = visit(links, node.FirstChild)
	}

	if node.NextSibling != nil {
		// A recursive call to next sibling
		links = visit(links, node.NextSibling)
	}

	return links
}

// getTimeout retries for the specified URL until the timeout
func getTimeout(url string, timeout time.Duration) (io.Reader, error) {
	endTime := time.Now().Add(timeout)

	for retry :=0 ; time.Now().Before(endTime) ; retry++ {
		resp, err := http.Get(url)
		if err != nil {
			time.Sleep(time.Second)
			fmt.Printf("Retry %d\n", retry)
			continue
		}
		return resp.Body, nil
	}

	return nil, fmt.Errorf("get %s: timed out after %f seconds", url, timeout.Seconds())
}