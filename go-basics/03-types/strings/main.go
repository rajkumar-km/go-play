/*
strings demonstrates the use of strings in Go
*/
package main

import "fmt"

func main() {
	// Demo string type
	fmt.Println("--- DemoString -------")
	DemoString()

	// Demo unicode
	fmt.Println("\n--- DemoUnicode -------")
	DemoUnicode()

	// Demo byte slices
	fmt.Println("\n--- DemoBytes -------")
	DemoBytes()

	// Demo string conversions
	fmt.Println("\n--- DemoStrconv -------")
	DemoStrconv()
}
