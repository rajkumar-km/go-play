/*
once demonstrates the Go sync.Once feature for lazy initialization

Lazy initializations are useful sometimes. If a program supposed to initialize big structure
but that may or may not be used depends on the usage. In such case, a lazy initialization
is effective.

If multiple goroutines accesses the data structure, then it must be guarded with locks.
Mutually exclusive locks for both read and write operations are costiler since we initialize
only once. sync.Once can be used here.

sync.Once consists of two values: a boolean flag and a mutex. The mutex guards both the boolean
and the client's data structure.
*/
package main

import (
	"fmt"
	"image"
	"sync"
	"time"
)

var (
	loadIconsOnce sync.Once
	icons         map[string]image.Image
)

// load icons initializes all the icons
func loadIcons() {
	// Initialize the icons here
	icons = map[string]image.Image{
		"logo": nil, // implement loadIcon("logo")
	}
	fmt.Println("Loading icons only once here")
}

// Icon returns the image by the tag
func Icon(name string) image.Image {
	// loadIconsOnce has a boolean flag which is set to false by default
	// This also has a mutex which will be locked before checking the flag
	// If the flag is false, then it performs the loadIcons() and releases the lock
	loadIconsOnce.Do(loadIcons)
	fmt.Println("Returning icon:", name)
	return icons[name]
}

func main() {
	go Icon("logo")
	go Icon("loading")
	go Icon("home")
	// wait for the goroutines to complete. use waitgroup in production.
	time.Sleep(100 * time.Millisecond)
}
