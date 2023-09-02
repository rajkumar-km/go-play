package main

import "fmt"

func main() {
	// Pointers in Go
	fmt.Println("--- DemoPointers -------")
	DemoPointers()

	// The new() function to create pointers
	fmt.Println("--- DemoNewFunction -------")
	DemoNewFunction()

	// Pointers as function arguments and return values
	fmt.Println("--- DemoPointerArgs -------")
	DemoPointerArgs()
}
