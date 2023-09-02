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

It is perfectly safe to return a pointer to a local variable from functions. The memory
is held as long as it has the reference.

Call by reference:
  - Passing pointer arguments to functions can affect the value of original variable

The new function:

  - new is another way to create a variable, but returns a pointer to the variable

  - This is just a syntactic convenience to create a pointer variable and not a fundamental notion

    var p *int = new(int)
    The above statement is equal to "var x int ; var p *int = &x"

  - Every call to new() function create a new instance of the variable

  - However there is a exception to this rule:

  - Call to new(struct{}) and new([0]type) can return the same address depends on the implementation.

  - Because struct{}, [0]int, [0]string, ... are considered zero bytes
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

	// Creating a Pointer using the built-in new() function
	// Allocates the Memory and returns the address
	// Go garbage collector is responsible to free the memory when it becomes inaccessible (say textptr = nil)
	{
		textptr := new(string)
		*textptr = "my string"
		fmt.Println(textptr, "=", *textptr)
		textptr = nil // Garbage collector to free the memory

		// Address of new(struct{}) can be same because it is considered as zero bytes
		t1 := new(struct{})
		t2 := new(struct{})
		fmt.Printf("Address of struct{} = %p, %p\n", t1, t2)

		// Address of zero size array can be same because it is considered as zero bytes
		t3 := new([0]string)
		t4 := new([0]string)
		fmt.Printf("Address of [0]string = %p, %p\n", t3, t4)
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

	// Call function returning a reference of local variable
	{
		p1 := localref()
		p2 := localref()
		fmt.Println("Each call creates a separate local variable, so pointer comparision is", p1 == p2)
		fmt.Println("Values stored in the reference are same:", *p1 == *p2)
	}

	// Call by reference
	{
		v := "hello"
		p := &v
		reverseInPlace(p)
		fmt.Println("hello after reverse:", v)
	}
}

// localref return the address of a local variable
func localref() *int {
	var x int = 100
	// It is perfectly safe to return the address of a local variable
	// Go held this in Heap memory as long as it has some reference
	return &x
}

// reverseInPlace reverses a string in the given address
func reverseInPlace(s *string) {
	b := []byte(*s) // this creates a copy of string
	for i, j := 0, len(b)-1; i < j ; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	// Update the reference to revesed string
	*s = string(b)
}
