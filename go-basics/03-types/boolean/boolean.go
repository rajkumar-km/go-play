/*
boolean demonstrates use of bool type in Go
*/
package main

import "fmt"

func main() {
	// Demo boolean type
	fmt.Println("--- DemoBool -------")
	DemoBool()
}

// DemoBool demonstrates the boolean type in Go.
// - The true and false values are booleans
// - Comparison operators like == and < produces bool value
// - Conditions on if and for statement are boolean
// - Unary ! operator can be used for bool negation
// - The operators && and || can be used to combine booleans
func DemoBool() {
	var x bool // false by default
	fmt.Printf("Default value of var x bool = %v\n", x)
	x = true

	// Comparison
	if x { // shorthand for x == true
		fmt.Println(`Use "if x" instead of "if x == true"`)
	}

	// Unary !
	x = false
	if !x {
		fmt.Println(`Use "if !x" instead of "if x == false"`)
	}

	// && has high precedene than ||
	var c byte = 'N'
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		fmt.Printf("%c is alphanumeric\n", c)
	}

	// No direct conversion to int type, so write a function
	fmt.Printf("No direct conversion to int type, so write a own function: %d\n", btoi(true))
}

// btoi converts bool value to int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
