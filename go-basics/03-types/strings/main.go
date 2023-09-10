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
	fmt.Println("--- DemoUnicode -------")
	DemoUnicode()

	// Demo byte slices
	fmt.Println("--- DemoBytes -------")
	DemoBytes()
}
