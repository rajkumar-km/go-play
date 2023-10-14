/*
error demonstrates the background of "error" interface in Go

  - error is actually an interface type which has the only method "Error() string"
  - So we can return any object in place of error after implementing Error() method
  - package "errors" contains a function "errors.New(msg)" which returns a interface
    error holding the pointer to simple object of struct "errorString"
*/
package main

import (
	"errors"
	"fmt"
)

const ENOENT = 2

func main() {
	// Create errors using errors.New()
	// Note that its internal implementation in such a way that every errors.New()
	// creates a new instance. It returns a pointer stored in interface rather than value.
	// func New(text string) error { return &errorString{text} }
	fmt.Println(errors.New("Invalid") == errors.New("Invalid")) // false

	err := systemExec()
	fmt.Println(err) // "error 2"

	var serr SystemError = err.(SystemError)
	// This would still print "error 2" since Error() is automatically invoked by fmt
	// even before trying String() method.
	fmt.Println(serr) // "error 2"

	fmt.Printf("%d\n", serr)    // 2
	fmt.Println(serr == ENOENT) // true
}

// systemExec is a model function that returns SystemError wrapper in error
func systemExec() error {
	return SystemError(ENOENT)
}

// SystemError is a custom error that satisfies the common error interface
// Note: Even a struct holding several information can also be a custom error
type SystemError int

func (se SystemError) Error() string {
	return fmt.Sprintf("error %d", int(se))
}
