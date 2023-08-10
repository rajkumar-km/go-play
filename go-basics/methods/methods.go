/*
Package main demonstates the use of methods in Go programming

Methods are just a function with a special receiver argument.
It can also be written as normal function by passing the object as argument

Why Method instead of Function?
- Methods provides the object oriented style
- Short and easy to understand
- Same function name can be reused for different types
- Methods can be called by value as well as pointer types.
  - Go automatically converts to appropriate type based on the method definition
  - obj->method() converts to (&obj) if the method expects pointer receiver
  - ptr->method() converts to (*ptr) if the method expects value receiver

Method definition restrictions:
- Method can be defined on the same package - the receiver type is defined
- It does not allow to define methods for built-in types or types defined in other packages.
*/
package main

import (
	"fmt"
	"strings"
)

// A Person is a struct that represents the user information
type Person struct {
	Id   int
	Name string
}

// Print is a method with value receiver: takes a copy of the receiver argument Person
// Exposed outside the current package as it starts in upper case
func (p Person) Print() {
	fmt.Println(p)
}

// update is a method with pointer receiver: this would affect the calling object
// Starting with lower case are not exported to access from another package
func (p *Person) update(name string) {
	p.Name = name
}

// MyString is a new type and stores the built in type string
type MyString string

// trim is a method defined on non-struct type
func (s MyString) trim() string {
	return strings.TrimSpace(string(s))
}

// main demonstates the use of methods in Go
func main() {
	p1 := Person{Id: 1, Name: "Alice"}
	p1.Print()
	// Go automatically passes the pointer as receiver (&p1)
	// So, a method with pointer receiver makes changes to original object
	p1.update("Alice Bob")
	p1.Print()

	var str MyString = " Hello "
	fmt.Println(str.trim())
}
