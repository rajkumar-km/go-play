/*
newfunction demonstrates the use of new() function in Go programming

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

// DemoNewFunction demonstrates using pointers in
// - Function arguments
// - Return values
func DemoNewFunction() {

	// Creating a Pointer using the built-in new() function
	// Allocates the Memory and returns the address
	// Go garbage collector is responsible to free the memory when it becomes inaccessible (say textptr = nil)
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
