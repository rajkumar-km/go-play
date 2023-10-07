/*
methods demonstates the use of methods in Go programming

Go is no exception while providing support to OOPs. But there is no accepted standard for
object-oriented programming.
- In Go, an object is simply a value or variable that has methods
- Methods are just a function with a special receiver argument associated with a type
- But, it can also be written as normal function by passing the object as argument

Why Method instead of Function?
- Methods provides the object oriented style
- Short and easy to understand
- Same function name can be reused for different types

Pointer vs Value receivers:
- Methods can be defined either as pointer type or value type receiver.
- Value receiver method takes a copy of an object
- Pointer receiver methods uses the original object and it can modify them

Calling methods with value vs pointer:
  - Go allows to call a method by both value and pointer irrespective of whether it is defined
    with pointer receiver or value receiver.
  - There are three possibilties:
    1. Both calling parameter and function receiver are same type
  - No issues. Go just invokes the method directly
    2. Calling parameter is pointer, but function receiver is value type
  - Go automatically deference the pointer (*ptr) and pass it to function
    3. Calling parameter is value type, but function receiver is pointer type
  - Go obtains the reference (address of the value) and pass it to the function
  - But here is the catch: The value should have the addressable memory. If it is stored in
    some variable or returned by new() function then it should have address. But, literals
    like Obj{} does not have the addressable memory and causes compile error.

- All put together:
  - Value receiver methods can be called by both value type and pointer type.
  - But, Pointer receiver methods can be called by pointers and addressable value types. So,
    this can not be called by values stored in interface/map that are not addressable.

- What is addressable value types? - Read more at https://go.dev/ref/spec#Address_operators

Method definition restrictions:
  - Method can be defined on the same package - the receiver type is defined
  - It does not allow to define methods for built-in types or types defined in other packages.
  - Although it allows methods for named types, it would not allow to define methods for named
    types that are themselves pointers. This is to avoid ambiguities since we already support
    value and pointer receiver methods.
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

// Traditional function (not object oriented method)
func Print(p Person) {
	fmt.Println(p)
}

// Print is a method with value receiver: takes a copy of the receiver argument Person
// Value receiver methods can be called by both value type and pointer type.
// Also, it is exposed outside the current package as it starts in upper case
// Go consider the name of this method as "Person.Print" and it does not conflict
func (p Person) Print() {
	fmt.Println(p)
}

// update is a method with pointer receiver: this would affect the calling object
// Pointer receiver methods can be called by pointers and addressable value types.
// Starting with lower case are not exported to access from another package
// Go consider the name of this method as "(*Person).update"
func (p *Person) update(name string) {
	p.Name = name
}

// MyString is a new type and internally stores the built in type string
// However, they both are treated as separate types and can not mix them.
// We can even add methods to the defined type MyString
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
	Print(p1)  // function call
	p1.Print() // method call equivalent to function call above

	// Go automatically passes the pointer as receiver (&p1)
	// So, a method with pointer receiver makes changes to original object
	p1.update("Alice Bob")
	p1.Print()

	// Also note that the pointer receiver functions won't work on non addressable values.
	// because there is no way to take the address of Person{}
	// Person{Id: 1, Name: "Alice"}.update("Alice Bob") // error: can not call pointer method update

	// This works becuase new() returns a addressable memory
	new(Person).update("Bob")

	var str MyString = " Hello "
	fmt.Println(str)
}
