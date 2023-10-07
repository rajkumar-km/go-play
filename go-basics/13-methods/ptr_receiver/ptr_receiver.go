/*
ptr_receiver demonstates using pointer receiver in methods

Since value receiver methods take a copy of the object, we can not modify its
properties. Say person.SetName("Bob") won't work since it works on separate copy.
A pointer receiver method is helpful here.

Pointer vs Value receivers:
- Methods can be defined either as pointer type or value type receiver.
- Value receiver method takes a copy of an object
- Pointer receiver methods uses the original object and it can modify them

Calling methods with value vs pointer:
  - Go allows to call a method by both value and pointer irrespective of whether it is defined
    with pointer receiver or value receiver.
  - There are three possibilties:

1. Both calling parameter and function receiver are same type:
  - No issues. Go just invokes the method directly
  - For pointer type, even it works if pointer contains nil value

2. Calling parameter is pointer, but function receiver is value type:
  - Go automatically deference the pointer (*ptr) and pass it to function
  - Panic occurs if the pointer is nil, because we can not dereference a nil pointer

3. Calling parameter is value type, but function receiver is pointer type:
  - Go obtains the reference (address of the value) and pass it to the function
  - But here is the catch: The value should have the addressable memory. If it is stored in
    some variable or returned by new() function then it should have address. But, literals
    like Obj{} does not have the addressable memory and causes compile error.

All put together:
  - Value receiver methods can be called by both value type and pointer type.
  - But, Pointer receiver methods can be called by pointers and addressable value types. So,
    this can not be called by values stored in interface/map that are not addressable.

What is addressable value types? - Read more at https://go.dev/ref/spec#Address_operators
*/
package main

import "fmt"

// A Person is a struct that represents the user information
type Person struct {
	Id   int
	Name string
}

// update is a method with pointer receiver: this would affect the calling object
// Pointer receiver methods can be called by pointers and addressable value types.
// Go consider the name of this method as "(*Person).update"
func (p *Person) Update(name string) {
	p.Name = name
}

// main demonstates the use of methods in Go
func main() {
	fmt.Println("Go automatically passes the address (&obj) while invoking a pointer receiver method")
	p1 := Person{Id: 1, Name: "Alice"}
	// Go automatically passes the pointer as receiver (&p1)
	// So, a method with pointer receiver makes changes to original object
	p1.Update("Alice Bob")

	// Note that the pointer receiver functions won't work on non addressable values.
	// because there is no way to take the address of Person{} in this case
	// Person{Id: 1, Name: "Alice"}.update("Alice Bob") // error: can not call pointer method update
	fmt.Println("Note that the pointer receiver functions won't work on non addressable values.")
	fmt.Println("Object literals that are not assigned to a variable are not addressable")

	// Values stored in the map are not addressable, because a map can grow/shrink its size
	// and it is free rearrange its elements. So, can not provide a fixed address of its elements.
	fmt.Println("Values stored in the map are not addressable")
	m := make(map[string]Person)
	m["Alice"] = Person{Id: 1, Name: "Alice"}
	// m["Alice"].Update("Bob") // compile error: cannot call pointer method Update

	// Values stored in array/slices are addressable
	fmt.Println("Values stored in the array/slices are addressable")
	s := []Person{
		{Name: "Alice"},
		{Name: "David"},
	}
	s[1].Update("Bob")

	// This works becuase new() returns a addressable memory
	fmt.Println("Values obtained by new() function are addressable")
	new(Person).Update("Bob")
}
