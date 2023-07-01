/*
Package functions demonstrates the functions in Go programming

Please note that Go does not provide overloading
Reason from Go:

	Method dispatch is simplified if it doesn't need to do type matching as well.
	Experience with other languages told us that having a variety of methods with the same name
	but different signatures was occasionally useful but that it could also be confusing and fragile in practice.
	Matching only by name and requiring consistency in the types was a major simplifying decision in Go's type system.
*/
package functions

import (
	"errors"
	"fmt"
)

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

// Play demonstrates the Go function calls
func Play() {
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
}
