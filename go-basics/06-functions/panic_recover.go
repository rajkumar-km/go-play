package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
)

// DemoPanicRecover demonstrates panic() and recover() calls in Go
//   - Go's type system catches most of the issues at compile time, but issues like out of slice
//     access, nil pointer access requires check at run time.
//   - Go's run time detects these mistakes and panics - Stops execution, runs all the deferred
//     calls and program crashes with a message and stack trace.
//   - Go’s panic mechanism runs the deferred functions before it unwinds the stack. It means we
//     can obtain the panic's trace within a deferred function using "runtime.Stack()" call.
//   - We can also use the built-in panic() function to cause panic during fatal errors. This
//     can be used to indicate a bug in the program. For example, a logical error where the control
//     is not expected reach at some point, or a missing developer's (not user's) configuration.
//   - Note: Errors like user's misconfiguration, incorrect input, failing IO are expected to be
//     handled gracefully without panic()
func DemoPanicRecover() {
	defer printStack()
	defer fmt.Println("DemoPanicRecover completed")
	for _, divisor := range []int{10, 1, 10} {
		// This would cause panic on third iteration while dividing by zero
		// panic: runtime error: integer divide by zero
		// This would execute all the defer calls from the stack and panics
		fmt.Printf("dividing 100 by %d\n", divisor)
		divide(100, divisor)
	}

	// MustCompile causes panic if regexp is invalid
	MustCompile("[abcd")
}

// divide performs the mathematical division
// Note that it causes panic when the divisor is zero
func divide(x, y int) int {
	return x / y
}

// printStack prints the current stack trace
// Go’s panic mechanism runs the deferred functions before it unwinds the stack
// So, a defer call to printStack() can be useful to debug the place of actual panic
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

// MustCompile compiles a regular expression and causes panic if it is invalid
// Note: MustCompile is already part of regexp package, This is just to demo panic()
func MustCompile(expr string) *regexp.Regexp {
	re, err := regexp.Compile(expr)
	if err != nil {
		panic(err) // panic accepts the argument of type "any"
	}
	return re
}
