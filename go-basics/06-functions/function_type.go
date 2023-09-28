package main

import "fmt"

// DemoFunctionTypes demonstrates functions as value types in Go
func DemoFunctionTypes() {
	var f func(int) int
	fmt.Printf("%v\n", f)
	
	p := add
	fmt.Printf("%d\n", p(10, 20))

	p = mul
	fmt.Printf("%d\n", p(10, 20))

	// compile error: cannot use pow2 (value of type func(x int) int) as func(x int, y int) int 
	// value in assignmentcompilerIncompatibleAssign
	// p = pow2

	// f is compatiple with type of pow2
	f = pow2
	fmt.Printf("%d\n", f(10))
}

func add(x int, y int) int {
	return x+y
}

func mul(x int, y int) int {
	return x*y
}

func pow2(x int) int {
	return x*x
}
