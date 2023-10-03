package main

import "fmt"

// Person identifies a person
type Person struct {
	Id int
}

// Member is an alias of Person
type Member = Person

// DemoTypeAlias demonstrates having type aliases in Go
//
//	type newName = name
//
// Type aliases are intended to be used during code repair to not to break clients.
// Type alias is just another name for the type. They can be used interchangably.
//
// Note: This is different from type declation which creates a new distinct type. They 
// can not be compared against other types. Also type declaration allows to add receiver
// methods even if it is created from a built-in types. But type alias is not capable of
// doing more than what the base type can do.
func DemoTypeAlias() {
	p := Person{1}
	q := Member{1}
	fmt.Println("Type alias can be used interchangably and mixed during comparisions: ", p == q)
}