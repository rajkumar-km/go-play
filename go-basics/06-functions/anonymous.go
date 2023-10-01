package main

import (
	"fmt"
	"strings"
)

// DemoAnonymous demonstrates anonymous functions in Go
//   - Named functions can be written in Go at the package level only
//   - In addition, Go supports function literals that can be written even inside another function.
//   - A function literal is written like a normal function declaration, but without function name.
//   - It is considered as Go expression and its value is called anonymous function.
//
// Why anonymous functions?
//   - Allows to define function at its point of use
//   - More importantly, those functions have access to their entire lexical environment
//   - So, a function value is not just action, but it can also have states.
//
// Closures:
//   - In programming languages, a closure means binding a function with its environment.
//   - For example a function referencing a variable outside its body.
//   - Since Go has closures, anonymous functions can still has references outside variables and
//     the scope of the variable is retained.
func DemoAnonymous() {
	// A simple anonymous function at the point of use
	source := "hello"
	transformed := strings.Map(func(r rune) rune { return r + 1 }, source)
	fmt.Printf("Transformed %q to %q\n", source, transformed)

	// Function closures that carries states
	f := newFibonacci() // creates variable f of type "func() int"
	fmt.Printf("Function closures to generate finonacci sequence:")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", f())
	}
	fmt.Println()

	// Anonymous functions can also be made recursive by separating the variable declaration
	// If you combine, variable declaration and assignment that it results in compile error
	// because the variable is not declared to use inside the function.
	var fact func(n int) int
	fact = func(n int) int {
		if n <= 0 {
			return 1
		}
		return n * fact(n-1)
	}
	fmt.Printf("Recursive anonymous function to find factorial(%d): %d\n", 4, fact(4))

	// Caveat: Capturing iteration variables
	// Sometimes, even expert programmers failed to understand closures
	cleanups := []func(){}
	for i := 0; i < 3; i++ {
		// i := i // Tip: enable this to define a local variable specific to this iteration
		// and it can be safely captured in anonymous functions.
		cleanups = append(cleanups, func() { fmt.Println("\tCleaning up index: ", i) })
	}
	fmt.Println("Caveat: capturing iteration variables results in accessing last value")
	for _, f := range cleanups {
		f()
	}
}

// newFibonacci returns a function that can be used to retrieve sequence of finonacci numbers
func newFibonacci() func() int {
	x := 0
	y := 1
	// Note that the returned function is a closure since it access the outer variables x and y
	// Even though the variables are out of scope after this return, the scope is retained since
	// it is bound to this anonymous function. So, this function carries states as well.
	return func() int {
		x, y = y, x+y
		return x
	}
}
