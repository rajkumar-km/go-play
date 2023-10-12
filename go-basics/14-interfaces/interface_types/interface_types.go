/*
interface_types demonstrates using interface as types in Go

How does an interface type work with concrete values?
  - Under the hood, an interface value can be thought of as a tuple consisting
    of a value and a concrete type: [value, type]
  - Example: [{65, 94}, main.Rectangle]
  - Since Go is a statically typed language, a type can not be a value. The internal representation
    is different and its hold only a type descriptor.
  - Dynamic dispatching: If an interface method is called, the corresponding method on the concrete
    type is called.
  - The compiler generates the code to retrieve the address of the method and invoke with the
    receiver argument as concrete value.

Assigning interface with pointer and value types:
  - Another point to remember is an interface can be assigned with both value and pointer type.
  - A copy is always made and stored in interface. So any changes to the object is not reflected
    on the other side.
  - In case of a pointer, only the pointer variable is copied. But the address it is pointing to
    is always same. So, storing a pointer type always reflects the changes outside.

Interface and the methods with pointer receiver:
  - If the interface is stored with a value type object, then it can not invoke the associated
    methods with pointer receiver.
  - Because, the value stored in the interface is not addressable, so it can not dereference
    it to pointer and invoke the method.
  - It works fine when the a copy of pointer variable is stored in interface. Because it can
    locate the actual value.

Interfaces comparability:
  - Interface can be compared if the concrete values stored in the interface are comparable
  - If the interface contains a slice, map, or function that it can not be comparable.
  - The default value of interface is nil and it can always be compared against nil
  - Mind that the interface is not nil if it stores a pointer to nil value. Because, interface
    actually holds a copy of pointer variable, and it can refer to nil internally.
*/
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Stringer is the interface that has String() method
type Stringer interface {
	String() string
}

// A simple Queue that supports Stringer interface to print its elements.
type Queue struct {
	elements []int
}

// String() returns the string representation of Queue
func (q *Queue) String() string {
	var buf bytes.Buffer
	for i := range q.elements {
		buf.WriteString(strconv.Itoa(q.elements[i]))
		buf.WriteByte(' ')
	}
	return buf.String()
}

// Enqueue inserts an element to Queue
func (q *Queue) Enqueue(v int) {
	q.elements = append(q.elements, v)
}

func main() {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)

	// [type descriptor, concrete value] = [nil, nil]
	var s Stringer
	fmt.Printf("%T\n", s) // fmt internally uses reflection to print its type descriptor
	// s.String() // runtime_error: nil pointer dereference

	// [type descriptor, concrete value] = [*Queue, &q]
	// A copy of pointer variable is stored in interface which still refers to original q
	s = &q
	fmt.Printf("%T\n", s)

	// Dynamic dispatching
	// s.String() => &q.String()
	qStr := s.String()
	fmt.Println(qStr)

	// [type descriptor, concrete value] = [*Queue, q]
	// A copy of original q is expected, but the copied value is not addressable
	// Tip: Converting String() to value receiver works for both value and pointer types
	// s = q // compile error: method String has pointer receiver

	// [type descriptor, concrete value] = [nil, nil]
	// Interface can be compared against nil
	s = nil
	fmt.Println(s == nil)

	// Interface is considered as not nil if it contains a "pointer to nil"
	// Because, interface actually holds a copy of pointer variable "*Queue"	
	var nilQ *Queue
	// [type descriptor, concrete value] = [*Queue, nilQ=nil]
	s = nilQ
	fmt.Println(s == nil)

	// Programmers can make this mistake easily
	// We pass "s" which is type of *Queue and value is nil
	Inspect(s)
}

// Inspect prints the Stringer's string value
func Inspect(s Stringer) {
	if s != nil {
		// Note that s is not nil if we pass a pointer to a nil value
		// s.String() is still invoked for *Queue since is a pointer receiver
		// Pointer receiver methods is invoked with nil, but it can cause panic
		// when we try to access any value elements inside.

		// fmt.Println(s.String())
	} else {
		fmt.Println("<nil>")
	}
}