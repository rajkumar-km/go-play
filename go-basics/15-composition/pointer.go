/*
pointer demonstrates composition using pointers in Go programming

The anonymous field we embed to the struct can be a pointer type.
  - Adding one more level of indirection, allows to reference a single embedded object
    to multiple derived objects.
  - So, this can create/modify relationship among different objects
  - All the properties and methods are derived indirectly through pointers.
*/
package main

import "fmt"

// Group indicates group of persons
type Group struct {
	Id    int
	Title string
}

// Print display the Group information
func (g Group) Print() {
	fmt.Println(g.Title)
}

// Person indicates a person
type Person struct {
	// Embed Group as pointer type
	// So, we can reference the single Group object for multiple Person with pointers
	*Group
	Id   int
	Name string
}

// DemoPtrToStructEmbed demonstrates the pointer to struct composition in Go
func DemoPtrToStructEmbed() {
	fmt.Println("A struct can be embedded to another as pointer type as well")

	p := Person{}

	// We have not initialized the *Group and it will be nil by default
	// Note that the Print() method is only defined in *Group
	// The following statement panics since it trying to dereference the value of
	// *Group to invoke the value receiver method Print().
	// Pointer receiver methods still works with nil pointer. Unless you have all the
	// methods defined as pointer receiver, it is safe to initialize the embedded object.

	// p.Print()

	g := Group{Title: "Maths"}
	p = Person{Name: "Alice", Group: &g}
	p2 := Person{Name: "Bob", Group: &g}
	p.Print()
	p2.Print()
}
