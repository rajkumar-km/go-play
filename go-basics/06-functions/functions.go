/*
function demonstrates the functions in Go programming
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

// main demonstrates the Go function calls
func main() {
	result := sum(10, 20, "10+20=")
	fmt.Println(result)

	// Blank identifer "_" to ignore function return values from the caller
	result2, _ := div(10, 5, "10/5=")
	fmt.Println(result2)

	// Error handling
	result2, err := div2(10, 0, "10/0=")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	fmt.Println(one(1, 0))
	fmt.Println(zero(1, 0))
}

// sum is a simple function
func sum(arg1 int, arg2 int, arg3 string) int {
	fmt.Print(arg3)
	result := arg1 + arg2
	return result
}

// div function supports multiple return values with error type
func div(arg1, arg2 int, arg3 string) (int, error) {
	result := 0
	var err error = nil
	fmt.Print(arg3)
	if arg2 == 0 {
		err = errors.New("can not divide by zero")
	} else {
		result = (arg1 / arg2)
	}
	return result, err
}

// div2 with named return values (all return values must be named, can not mix)
// Go declares the variables for return values
func div2(arg1, arg2 int, arg3 string) (result int, err error) {
	fmt.Print(arg3)
	if arg2 == 0 {
		err = errors.New("can not divide by zero")
	} else {
		result = (arg1 / arg2)
	}
	return // Optional to provide return values. Plain return statement automatically returns the named return values
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
