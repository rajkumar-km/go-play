/*
Package main demonstrates the use of pointers in Go programming
A pointer is a special type of variable that actual holds the memory address of
an another variable.
*/
package main

import "fmt"

// main demonstrates the use of pointers in Go programming
// - Declaring and assigning values to pointer
// - Pointer to pointer
// - No pointer arithmetic in Go, but pointers comparisions allowed
func main() {
	// Declaring a Pointer of type T
	// Syntax: var p *T
	var p *int     // default to nil
	fmt.Println(p) // prints <nil>

	// Initializing/Dereferencing a Pointer
	var number int = 10000
	p = &number
	fmt.Println("var number int =", number, ", addr =", &number)
	fmt.Println("var p *int =", p, ", dereferencing *p =", *p)

	// Creating a Pointer using the built-in new() function
	// Allocates the Memory and returns the address
	// Go garbage collector is responsible to free the memory when it becomes inaccessible (say textptr = nil)
	textptr := new(string)
	*textptr = "my string"
	fmt.Println(textptr, "=", *textptr)
	textptr = nil // Garbage collector to free the memory

	// Pointer to Pointer
	pp := &p
	fmt.Println("p =", p, "*p =", *p)
	fmt.Println("pp =", pp, "*pp =", *pp, "**pp =", **pp)

	// Like C/C++, No pointer arithmetic in Go
	// An equality comparision is allowed however
	p2 := p
	if p == p2 {
		fmt.Println("Pointer p and p2 are equal")
	}
}
