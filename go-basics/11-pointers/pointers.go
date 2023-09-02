/*
pointers demonstrates the use of pointers in Go programming

# A variable is a piece of storage containing a value

A pointer is a special type of variable that holds the memory address of an another variable.

	// variable of type "pointer to int". All pointer types by default initialized to nil
	var p *int

	// &x denotes the address of variable x
	// Variables are sometimes described as addressable values
	// pointer p points to x
	var x int
	p = &x

	// update/access the values with *
	*p = 10
	var y int

	// comparision works for the address and not the values
	fmt.Println(&x == &x, &x == &y, &y == nil) // "true false false"
*/
package main

import "fmt"

// DemoPointers demonstrates the use of pointers in Go programming
// - Declaring and assigning values to pointer
// - Pointer to pointer
// - No pointer arithmetic in Go, but pointers comparisions allowed
func DemoPointers() {
	// Declaring a Pointer of type T
	// Syntax: var p *T
	{
		var p *int     // default to nil
		fmt.Println(p) // prints <nil>
	}

	// Initializing/Dereferencing a Pointer
	{
		var number int = 10000
		var p *int = &number
		fmt.Println("var number int =", number, ", addr =", &number)
		fmt.Println("var p *int =", p, ", dereferencing *p =", *p)
	}

	// Pointer to Pointer
	{
		p := new(int)
		pp := &p
		fmt.Println("p =", p, "*p =", *p)
		fmt.Println("pp =", pp, "*pp =", *pp, "**pp =", **pp)
	}

	// Pointer comparisions
	// Like C/C++, No pointer arithmetic in Go
	// An equality comparision is allowed however
	{
		var x int
		p1, p2 := &x, &x
		fmt.Println("p1 and p2 points to x, so p1 == p2:", p1 == p2)
	}
}
