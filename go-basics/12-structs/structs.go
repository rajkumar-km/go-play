/*
Package main demonstrates the use of structures in Go programming

A struct is a user-defined type that contains a collection of named fields
It is like a light weight class without inheritance (while supporting composition)

Exported vs Unexported Structs and Struct Fields:
- As usual the stucts starting with Upper case letter is exported to another package
- The fields inside the structs also has restrictions
- Fields starting with upper case letters are exported, and not the lower case letters.
- The structs/fields starting with lower case letter is accessible only within the package

Structs are value types:
- Performs a new copy during an assignment or passing to a function

Struct Equality:
- Two struct variables are equal if all their corresponding fields are equal -
*/
package main

import "fmt"

// A Person represents the user information
type Person struct {
	id                  int
	age                 int
	firstName, lastName string
}

// main demonstrates the use of structs in Go
// - Struct declaration and initialization
// - Accessing the fields
// - Pointer to struct
// - Using new() function to create an instance of struct
func main() {
	// Declaring and Initializing a struct
	var person Person
	person.lastName = "Alice"
	fmt.Println(person)

	person1 := Person{} // All set to default values
	fmt.Println(person1)

	person2 := Person{id: 1, lastName: "Alice"} // default: age=0, firstName=""
	fmt.Println(person2)

	person3 := Person{
		id:        2,
		lastName:  "Bob",
		age:       33,
		firstName: "Lyn", // trailing comma is mandatory when breaking down as seperate lines
	}
	fmt.Println(person3)

	// Accessing fields of a struct
	person.firstName = "Alice"
	fmt.Println(person.firstName)

	// Pointer to a struct
	person4 := &Person{id: 4} // or person4 := Person{} and ptr := &person4
	fmt.Println(person4)
	fmt.Println((*person4).id)
	fmt.Println(person4.id) // Same as above, Go automatically dereference the pointer and access the field

	// Using the new() function
	person5 := new(Person)
	person5.id = 5
	fmt.Println(person5)
}
