/*
nil_receiver demonstrates calling methods with nil pointer of corresponding type

Make sure to add the documentation for the types that allows nil as valid value
*/
package main

import "fmt"

type Person struct {
	Id int
}

// PersonList contains list of Person
// A nil *PersonList indicates an empty list
type PersonList []Person

// Print prints the persons list (the ids)
// Prints as "empty list" if *PersonList is nil
func (pl *PersonList) Print() {
	if pl == nil {
		fmt.Println("empty list")
	} else {
		for _, p := range *pl {
			fmt.Println(p.Id)
		}
	}
}

// Print2 prints the list but it requires value type
func (pl PersonList) Print2() {
	for _, p := range pl {
		fmt.Println(p.Id)
	}
}

func main() {
	var persons *PersonList = nil
	// Note that the method is invoked even with nil value since it uses pointer receiver
	persons.Print()

	// persons.Print2() // panic: runtime error: invalid memory address or nil pointer dereference

	// Here, we use value type, but the internal value []Person is nil
	// Since we have the instance of PersonList, we can invoke the method
	var persons2 PersonList = nil
	persons2.Print2()
}
