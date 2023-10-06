/*
functions demontrates using functions in Go and handling errors
*/
package main

import "fmt"

func main() {

	// Demo functions
	fmt.Println("--- DemoFunctions -------")
	DemoFunctions()

	// Errors
	fmt.Println("--- DemoErrors -------")
	DemoErrors()

	// Function Types
	fmt.Println("--- DemoFunctionTypes -------")
	DemoFunctionTypes()

	// Anonymous functions
	fmt.Println("--- DemoAnonymous -------")
	DemoAnonymous()

	// Variadic functions
	fmt.Println("--- DemoVariadic -------")
	DemoVariadic()

	// Defer
	fmt.Println("--- DemoDefer -------")
	DemoDefer()

	// Panic and Recover
	fmt.Println("--- DemoPanicRecover -------")
	DemoPanicRecover()
}
