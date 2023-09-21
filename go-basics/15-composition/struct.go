/*
composition demonstrates composition in Go programming

Object composition:
  - Go does not support inheritance, but this can be achieved through composition
  - A struct can be embedded inside another struct by simply specifying the name
  - All the properties and functions are delegated to the outer struct
  - Go builds a table of pointers for each methods. Methods from the inner
    struct are promoted to outer struct if they don't exist
  - This table is used to identify if a struct implements all the methods in
    an inteface
*/
package main

import "fmt"

// A CommonObject provides common fields for any object type
type CommonObject struct {
	Id   int
	Name string
}

// Print display the information of CommonObject
func (c *CommonObject) Print() {
	fmt.Println("\tCommon: Id =", c.Id, ", Name =", c.Name)
}

// A Folder object that inherits CommonObject
// - Specify the name of a struct inside another struct for composition
type Folder struct {
	Level        int
	CommonObject // composition -- Folder inheriting all properties of CommonObject
}

// Print display the information of Folder object
// This also overrides the Print method in CommonObject
// But, it can invoke inner struct's Print with the full name
func (f *Folder) Print() {
	f.CommonObject.Print()
	fmt.Println("\tFolder: Level =", f.Level)
}

// A Volume object inherting *Folder object (as reference type)
// Composition with reference type requires initialization
type Volume struct {
	Size int
	*Folder
}

// DemoStructEmbed demonstrates the struct composition in Go
func DemoStructEmbed() {
	fmt.Println("A struct can be embedded to another for inheritance")

	// Initializing a folder
	// Literals must use the full form to initialize embedded structs
	var folder Folder = Folder{Level: 1, CommonObject: CommonObject{Id: 1, Name: "folder"}}
	folder.Print()

	// Initialing volume
	var vol1 Volume = Volume{Size: 100, Folder: &folder}
	// Attributes can be accessed with shorthand names
	fmt.Println("\tInherited attributes can be accessed directly:", vol1.Id, vol1.Name)
	// Attributes can also be accessed like 'vol1.CommonObject.Name'
	// Not the best practice, but useful in case if the original and embedded type has the same attribute.

	// Literals must provide the full path
	var vol2 Volume = Volume{Size: 200, Folder: &Folder{Level: 2, CommonObject: CommonObject{Id: 2, Name: "folder2"}}}
	fmt.Println("\tUnfortunately no shorthand available for literals:", vol2.Id, vol2.Name)
}
