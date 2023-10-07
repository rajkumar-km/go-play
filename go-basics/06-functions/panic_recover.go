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
// Panic:
//   - Go's run time detects these mistakes and panics - Stops execution, runs all the deferred
//     calls and program crashes with a message and stack trace.
//   - Go’s panic mechanism runs the deferred functions before it unwinds the stack. It means we
//     can obtain the panic's trace within a deferred function using "runtime.Stack()" call.
//   - We can also use the built-in panic() function to cause panic during fatal errors. This
//     can be used to indicate a bug in the program. For example, a logical error where the control
//     is not expected reach at some point, or a missing developer's (not user's) configuration.
//   - Note: Errors like user's misconfiguration, incorrect input, failing IO are expected to be
//     handled gracefully without panic()
// Recover:
//   - Giving up is usually the best way to treat the panic since it indicates a bug.
//   - But, sometimes we may want to recover from the situtation. Say a web server want to close
//     the client connection when it encounters a unexpected problem.
//   - A recover() function needs be called inside a deferred function. In such case, if the
//     enclosing function encounters panic(), as usual the deferred functions are executed and
//     the recover() call inside the deferred function catches the panic() error for processing
//     and returns normally without panic.
//   - Note that the function left the execution at the point of panic and can not execute the
//     remaining statements.
//   - If recover() is called without any panic then it simply returns nil.
// Recovering from a panic without careful judgement is a bad practise:
//   - Because the state of the package's variable is unknown
//   - It could have incomplete update to data structures, pending locks, unclosed connections or
//     file descriptors.
//   - More importantly, the bugs get unnoticed by just logging the details to a file.
func DemoPanicRecover() {
	divideByZero()
	invalidRegexp()
}

// divideByZero demonstrates panic caused when divide by zero
func divideByZero() {
	defer fmt.Println("divideByZero completed")
	defer recoverFromErrors() // recover()

	for _, divisor := range []int{10, 0} {
		// This would cause panic on second iteration while dividing by zero
		// This would execute all the defer calls from the stack and panics
		fmt.Printf("dividing 100 by %d\n", divisor)
		fmt.Printf("%d\n", 100/divisor) // panic: runtime error: integer divide by zero
	}
}

// recoverFromErrors is intended to be a deferred function to recover from panics
func recoverFromErrors() {
	p := recover()
	if p != nil {
		fmt.Println("Function encountered a panic:")
		printStack()
	}
}

// printStack prints the current stack trace
// Go’s panic mechanism runs the deferred functions before it unwinds the stack
// So, a defer call to printStack() can be useful to debug the place of actual panic
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

// invalidRegexp demonstrates invoking panic()
func invalidRegexp() (retErr error) {
	defer func() {
		if p := recover(); p != nil {
			// Prevents the function from panic and modify its return value as error
			retErr = fmt.Errorf("internal error: %v", p)			
		}
	}()

	// MustCompile causes panic if regexp is invalid
	MustCompile("[abcd")
	return nil
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
