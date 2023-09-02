/*
pointerargs demonstrates the use of pointers in function arguments and return values

  - It is perfectly safe to return a pointer to a local variable from functions. The memory
    is held as long as it has the reference.
  - Call by reference: Passing pointer arguments to functions can affect the value of
    original variable
*/
package main

import "fmt"

// DemoPointerArgs demonstrates the use of pointers as arguments
func DemoPointerArgs() {

	// Call function returning a reference of local variable
	p1 := localref()
	p2 := localref()
	fmt.Println("Each call creates a separate local variable, so pointer comparision is", p1 == p2)
	fmt.Println("Values stored in the reference are same:", *p1 == *p2)

	// Call by reference
	v := "hello"
	p := &v
	reverseInPlace(p)
	fmt.Println("hello after reverse:", v)
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
