package main

import "fmt"

// main function is the starting point of execution
// Demonstrates the Go types, type assertions, and type conversions.
func main() {
	// Demo types
	fmt.Println("--- DemoTypes -------")
	DemoTypes()

	// Type convertions
	fmt.Println("--- DemoTypeConversion -------")
	DemoTypeConversion()

	// Named Type Declarations
	fmt.Println("--- DemoTypeDeclaration -------")
	DemoTypeDeclaration()

	// Type aliases
	fmt.Println("--- DemoTypeAlias -------")
	DemoTypeAlias()

	// Type assertions
	fmt.Println("--- DemoTypeAssertion -------")
	DemoTypeAssertion()
}
