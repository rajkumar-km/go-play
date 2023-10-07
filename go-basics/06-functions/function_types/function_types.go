/*
function_types demonstrates functions as value types in Go
Functions are first-class values in Go and it carries both data and action.
This can be used like any other values:
- Store in a variable with function types.
- Pass it to functions, or return from functions
- In addition, you can invoke function stored in a variable

Example: var f func(int) int
- f is the variable with type "func(int) int"
- By default initialized to nil, so calling f() would panic
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var f func(int) int
	if f == nil {
		fmt.Println("Function variables are nil by default. They can be compared only against nil")
		fmt.Println("Function values are not comparable, so it can not be map keys")
	}

	// f is compatiple with type of pow2 "func(int) int"
	f = pow2
	fmt.Printf("Function variables can be assigned with defined functions and invoked: %d\n", f(10))

	p := add
	fmt.Printf("Shorthand assignment auto initialize function variables: %d\n", p(10, 20))

	p = mul
	fmt.Printf("Function variables once declared can be assigned with exactly same type of functions: %d\n", p(10, 20))

	// compile error: cannot use pow2 (value of type func(x int) int) as func(x int, y int) int
	// value in assignmentcompilerIncompatibleAssign
	// p = pow2

	// f can be assigned with an anonymous function as well if that is compatiple
	f = func(x int) int {
		return int(math.Pow(float64(x), 2))
	}
	fmt.Printf("Function variables can also be assigned with anonymous function: %d\n", f(10))
}

func add(x int, y int) int {
	return x + y
}

func mul(x int, y int) int {
	return x * y
}

func pow2(x int) int {
	return x * x
}
