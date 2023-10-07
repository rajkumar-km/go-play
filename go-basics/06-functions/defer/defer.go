/*
defer demonstrates using defer function calls in Go

	defer <functioncall()>
	- The defer keyword evaluates the function call expression, but defer the call until the
	completion of enclosing function.
	- Defer call is useful for the following operations.
	1. Closing files
	2. Unlock mutex
	3. Other cleanup tasks
	- We need to have the defer calls immediately after opening a file or locking a mutex.
	This gets automatically called at the end. Defer calls are executed even if panic() occurs.
	- But, note that the resource is released only at the end of the function. Sometimes you might
	want to free the resources before. Either you can perform close without defer or split it in
	a separate function to use defer.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	// Use defer to audit a functionality
	// Not that audit() is not the deferred call here
	// A func value returned by audit() is registred as defer call.
	defer audit("main")() // Note the extra parentheses to invoke returned func value

	// wget example for defer
	url := "https://google.com"
	content, err := wget(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save it in a file
	fName, err := save(content)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Saved in %s\n", fName)
	}

	// Lanuch a goroutine to cache the results
	go cache(url, fName)
}

// wget retrieve the given url and save the result in systems temp dir as "index.<random>.html"
func wget(url string) (content []byte, retErr error) {
	// Get the URL content
	resp, retErr := http.Get(url)
	if retErr != nil {
		return
	}
	defer resp.Body.Close() // deferred to the end of enclosing function

	content, retErr = io.ReadAll(resp.Body)
	if retErr != nil {
		return
	}

	return
}

// save writes the content to a tempfile and returns the name
func save(content []byte) (retOutFile string, retErr error) {
	// Save it in a file
	file, retRrr := os.CreateTemp(os.TempDir(), "index.*.html")
	if retRrr != nil {
		return
	}

	defer func() {
		// This could have been written like "defer file.Close()"
		// But, some file system operations (such as NFS) does not throw any error on os.Write
		// But returns error while os.Close() is invoked. So, ignoring those errors can cause
		// serious data loss
		err := file.Close()
		if err != nil {
			// Anonymous functions has reference to function's local variables.
			// So, a deferred function here can modify the functions return value
			retErr = err
		} else {
			// Note that the return statement automatically updates the result variables and
			// this can be accessed inside the deferred anonymous function
			fmt.Printf("Returning filename: %s\n", retOutFile)
		}
	}() // Note the function call "()". This is required

	_, retRrr = file.Write(content)
	if retRrr != nil {
		return
	}

	// Update the return value
	return file.Name(), nil
}

// wgetCache is a map of URL and corresponding output file
var wgetCache = make(map[string]string)

// wgetMu is a mutex to synchronize access to wgetCache
var wgetMu sync.Mutex

// cache caches the wget output
// defer example for mutex unlock
func cache(url, fName string) {
	wgetMu.Lock()
	// Include a defer statement for unlock immediately after Lock()
	// This would for sure unlock the mutex even if panic() occurs.
	// No more worries about unlocking the mutex before every return
	defer wgetMu.Unlock()
	wgetCache[url] = fName
}

// audit is intended to be used in the defer call
// It prints the enter Audit record and returns a func value that can print the
// exit Audit record.
func audit(msg string) func() {
	start := time.Now()
	log.Printf("Audit enter: %s", msg)
	return func() { log.Printf("Audit exit: %s (%s)", msg, time.Since(start)) }
}
