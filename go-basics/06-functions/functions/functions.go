/*
functions demonstrates the functions in Go programming
  - function let us wrap the sequence of statements that can be reused
    func name(parameter-list) (result-list) {
    body
    }
  - parameter-list - function parameters which are local variables to function
  - result-list - return value parameters. parenthesis can be ignored if zero or one unnamed return
    value.
  - body contains the statements. Sometimes the function is just declared without { body }. This
    means the function is implemented in other language other than Go.

Go functions work as call-by-value
  - A copy of all the input parameters are available inside the function. So the changes does
    not affect the original variable.
  - However, if you pass pointers, slice, map, chan, or function, then it affects the original
    value. Still it is call-by-value, because it copies the small structure holding the pointer
    or other references types. But the structure internall points to the same memory location
    that affects the original value.

Go does not provide overloading. Here is the reason from Go:
  - Method dispatch is simplified if it doesn't need to do type matching as well.
  - Experience with other languages told us that having a variety of methods with the same name
    but different signatures was occasionally useful but that it could also be confusing and
    fragile in practice.
  - Matching only by name and requiring consistency in the types was a major simplifying decision
    in Go's type system.
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Simple function
	fmt.Println("Simple function")
	res := sum(10, 5)
	fmt.Println("\tfunc sum(x int, y int) int:")
	fmt.Println("\tsum(10, 5) =", res)

	// Function returning multiple values
	fmt.Println("Function with multiple return values")
	res2, _ := div(10, 5) // Use blank identifer "_" to ignore return values from the caller
	fmt.Println("\tfunc div(x, y int) (int,error)")
	fmt.Println("\tdiv(10, 5) =", res2)

	// Function with named return values
	fmt.Println("Function with named return value parameters")
	res3 := sub(10, 5)
	fmt.Println("\tfunc sub(x, y int) (result int)")
	fmt.Println("\tsub(10, 5) =", res3)

	// Function with unnamed parameters
	fmt.Println("Function with unnamed input parameters")
	fmt.Println("\tfunc one(int, _ int) int:", one(1, 0))
	fmt.Println("\tfunc zero(int, int) int:", zero(1, 0))

	// Recursive function
	fmt.Println("Recursive function")
	fmt.Printf("\tfactorial(%d)= %d\n", 5, factorial(5))
}

// sum add two integers and returns the sum
func sum(x int, y int) int {
	result := x + y
	return result
}

// div performs division operation and returns the result
// Note that we can combine input parameters having same type. Don't need to repeat the type
// Also note that the function has multiple return values with error type
func div(x, y int) (int, error) {
	result := 0
	var err error = nil
	if y == 0 {
		err = errors.New("can not divide by zero")
	} else {
		result = (x / y)
	}
	return result, err
}

// sub perform subtraction and returns the result
// Note the named return value parameters (all return values must be named, can not mix)
// Go declares the variables for return values
func sub(x, y int) (result int) {
	result = (x - y)
	// Return values are optional, so an empty return statement is enough.
	// Go automatically returns the named return values. But use this sparingly since it is
	// hard to understand when having multiple return values
	return // bare return
}

// Function with blank identifer as parameters
// Sometimes we need to implement a function with certain arguments to satisfy the common
// interface. But the implementation does not require all the parameters. So, we can use
// blank identifier to indicate compiler as well as other developers that the argument is
// not used in the function
//
// Secondly, an unnamed parameter (without blank identifier) is generally assigned with the same
// as the data type
func one(int, _ int) int {
	// int becomes "var int int"
	return int
}

// But, here both the variables are int, so we can not access them
func zero(int, int) int {
	return 0
}

// factorial returns the factorial number of n
// n! = n * (n - 1) * (n - 2) * ... * 1
// Note that it uses recursive calls for calculation
func factorial(n uint) uint {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
