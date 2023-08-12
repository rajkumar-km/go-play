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
  - obj.method() converts to (&obj) if the method expects pointer receiver
  - ptr.method() converts to (*ptr) if the method expects value receiver

- However, there is a difference between value vs pointer receiver
  - Value receiver methods can be called by both value type and pointer type.
  - Pointer receiver methods can be called by pointers and addressable value types. So,
    this can not be called by values stored in interface/map.

- What is addressable value types?
  - Read more at https://go.dev/ref/spec#Address_operators

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
// Value receiver methods can be called by both value type and pointer type.
// Exposed outside the current package as it starts in upper case
func (p Person) Print() {
	fmt.Println(p)
}

// update is a method with pointer receiver: this would affect the calling object
// Pointer receiver methods can be called by pointers and addressable value types.
// Starting with lower case are not exported to access from another package
func (p *Person) update(name string) {
	p.Name = name
}

// MyString is a new type and internally stores the built in type string
// However, they both are treated as separate types and can not mix them.
type MyString string

// String returns the trimmed string after removing leading/trailing spaces.
// So, methods can be defined for any types other than pointers/interfaces.
// This implements the String() method in fmt.Stringer interface, so that
// it is invoked automatically while passing to fmt.Println().
// Note that this won't be invoked if you use a pointer receiver. Because
// value types stored in an interface isn't addressable.
func (s MyString) String() string {
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

	// Also note that the pointer receiver functions won't work on
	// non addressable values.
	// error: can not call pointer method update
	//   Person{Id: 1, Name: "Alice"}.update("Alice Bob")
	// error: method update has pointer receiver
	//   interface{update(name string)}(p1)
	var str MyString = " Hello "
	fmt.Println(str)
}
