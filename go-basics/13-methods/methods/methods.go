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

Methods and named types:
  - Go does not allow to add methods on built-in types, arrays, slices, maps, channels, and
    functions.
  - But, we can indirectly add methods by defining a named type referencing built-in types.
  - Example:
  - type MyList []int
  - func (l *MyList) Sum() {}
  - Another example from net/http that has named type for a function and add methods to a function
  - type HandlerFunc func(w ResponseWriter, r *Request)
  - func (f *HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w,r) }

Method definition restrictions:
  - Method can be defined on the same package - the receiver type is defined
  - It does not allow to define methods for built-in types or types defined in other packages.
  - Although it allows methods for named types, it would not allow to define methods for named
    types that are themselves pointers. This is to avoid ambiguities since it supports
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
// Go consider the name of this method as "Person.Print" and it does not conflict
func (p Person) Print() {
	fmt.Println(p)
}

// MyString is a new type and internally stores the built in type string
// However, they both are treated as separate types and can not mix them.
// We can even add methods to the defined type MyString
type MyString string

// String returns the trimmed string after removing leading/trailing spaces.
// So, methods can be defined for any named types other than pointers/interfaces.
// This implements the String() method of fmt.Stringer interface, so that
// it is invoked automatically while passing it to fmt.Println().
//
// Note that this won't be invoked if you use a pointer receiver. Because
// value types stored in an interface isn't addressable. See ptr_receiver
// example for more details.
func (s MyString) String() string {
	return strings.TrimSpace(string(s))
}

func main() {
	p1 := Person{Id: 1, Name: "Alice"}
	Print(p1)  // function call
	p1.Print() // method call equivalent to function call above

	var str MyString = " Hello "
	fmt.Println(str) // fmt.Println internally invokes str.String()
}
