/*
struct demonstrates the use of structures in Go programming

A struct is a user-defined type that contains a collection of named fields of
different types. It is like a light weight class without inheritance
(while supporting composition)

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

import (
	"bufio"
	"fmt"
	"strings"
)

// A Person represents the user information
type Person struct {
	id                  int
	age                 int
	firstName, lastName string
}

func main() {
	createStructs()
	pointerToStructs()
	emptyStruct()
	structConvertCompare()
	// Also, look at 15-composition/struct.go to learn about inheritance by embedding structs
}

// createStructs demonstrates:
// - Struct declaration and initialization
// - Accessing the fields
func createStructs() {
	// Declaring and Initializing a struct is simple
	// Since it is a value type, it is initialized by default to zero values
	var person Person
	fmt.Println(`var person Person // auto initializes the struct since it is a value type`)
	fmt.Printf("\t%#v\n", person)

	// Struct can also be created using struct literal
	person1 := Person{} // All set to default values
	fmt.Println(`person1 := Person{} // empty struct literal`)
	fmt.Printf("\t%#v\n", person1)

	person2 := Person{id: 1, lastName: "Alice"} // default: age=0, firstName=""
	fmt.Println(`person2 := Person{id: 1, lastName: "Alice"} // struct literal`)
	fmt.Printf("\t%#v\n", person2)

	// The short form of struct literal used only when:
	// - Structs that has fixed fields and not expected to change (eg, Point(x, y), RGBA(r,g,b,a))
	// - Within the package if required. In case if the struct has unexported fields, then this
	//   form can not be used from other packages.
	// - The order matters when specifying values
	// - All the values must be provided and we can not omit fields unlike the other form
	person3 := Person{3, 35, "joe", "carter"}
	fmt.Println(`person3 := Person{3, 35, "joe", "carter"} // all fields must be specified and in order`)
	fmt.Printf("\t%#v\n", person3)

	// Note: A trailing comma is mandatory when breaking down as seperate lines
	person4 := Person{
		id:        2,
		lastName:  "Bob",
		age:       33,
		firstName: "Lyn", // trailing comma
	}

	// Accessing fields of a struct
	person4.firstName = "Alice"
	fmt.Printf("Struct fields can be accessed like: person.firstNam\n")
	fmt.Printf("\tperson.firstName = \"Alice\"\n")
	fmt.Printf("\tperson.firstName // returns %q\n", person3.firstName)
}

// pointerToStructs demonstrates using a pointer to struct
// and using the new() function to create an instance of struct
func pointerToStructs() {
	// Pointer to a struct
	// Using &T{..} is equivalent to new(T). This would create a variable of type T
	// and returns the address. So this can even be passed to functions directly
	// without assigning to a pointer variable.
	person := &Person{id: 1} // or person := Person{} and ptr := &person
	fmt.Println(`Pointers and structs`)
	fmt.Printf("\tperson := %#v\n", person)
	fmt.Printf("\t(*person).id = %d\n", (*person).id)
	// Same as above, Go automatically dereference the pointer and access the field
	fmt.Printf("\tperson.id = %d // go automatically dereference the pointer\n", person.id)

	// Note:  The assignment is not valid if the left hand side is not variable
	// (Person{}).age = 10 // compile error: UnassignableOperand
	(&Person{}).age = 10 // valid when assigning to pointer

	// Unlike map values, struct fields are addressable
	p := &person.id
	*p = 4

	// Using the new() function
	// The new() function is just a helper to allocate the memory of type T and zero it. This
	// can be used to avoid having a temporary variable and taking the address of it.
	// Note: Do not confuse it with make(). make only works with slice, map, and chan type and
	// performs the initialization. Also, make does not return a pointer.
	person5 := new(Person)
	fmt.Printf("The new(Person) returns a pointer to: %#v\n", person5)

	// A struct can not contain a field with type of itself, but it can have pointer so that
	// a linked list or tree can be implemented
	type BinaryTree struct {
		val         int
		left, right *BinaryTree
	}
}

// emptyStruct demonstrate the use of empty struct{} in Go
// A struct with no fields is considered zero size. So this can be used in some places
// instead of bool to save space.
func emptyStruct() {
	fmt.Println("Map can be used like sets: map[string]struct{}")
	in := "hi this is for testing words and it removes the the duplicates words"
	fmt.Print("\tBefore: ", in, "\n")
	fmt.Print("\tAfter :")

	r := strings.NewReader(in)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	set := make(map[string]struct{})
	for scanner.Scan() {
		word := scanner.Text()
		if _, ok := set[word]; !ok {
			set[word] = struct{}{}
			fmt.Print(" ", word)
		}
	}

	fmt.Println() // prints a newline
}

// structConvertCompare demonstrates converting and comparing one struct to another
func structConvertCompare() {
	type restPerson struct {
		ID                  int `json:"id"`
		firstName, lastName string
	}

	type grpcPerson struct {
		ID int `json:"ID"`
		// Note: the order must be same for conversion
		// Convertion would fail If we declare like "lastName, firstName string"
		firstName, lastName string
	}

	rjoe := restPerson{ID: 1, firstName: "Joe", lastName: "Carter"}
	gjoe := grpcPerson(rjoe)
	fmt.Println(`Struct objects can be converted to one type another provided:
	- All the field name and types are equal
	- The fields are declared in the exact same order
	- Tags does not matter from go 1.8`)

	expected := grpcPerson{ID: 1, firstName: "Joe", lastName: "Carter"}
	fmt.Println(`Struct objects can be compared if:
	- all the fields in the struct are comparable
	- So, a comparable struct can even be used as "key" in a map`)
	fmt.Println("\tcompare status: ", gjoe == expected)
}
