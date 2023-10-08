/*
composition demonstrates composition in Go programming

Object composition:
  - Go does not support inheritance, but this can be achieved through composition
  - A struct can be embedded inside another struct by simply specifying the name
  - All the properties and methods are delegated to the outer struct
  - Go builds a table of pointers for each methods. Methods from the inner
    struct are promoted to outer struct if they don't exist
  - This table is also used to identify if a struct implements all the methods in
    an interface

Compiler resolves a selector by:
  - First look at the direct method in the type
  - Second, Methods promoted from one level embedded types. If multiple methods are
    promoted from the same level, then it raises a compile error for ambiguity. it does
    not report error as long as the method is not called by derived object.
  - Next, looks for any methods promoted third level of embedding and so on.

Derived object and parent object are considered two distinct types:
  - You can not pass a derived object as an argument to the parent type like
    other object oriented languages.
  - Although the derived object has the embedded parent type, we must explicitly
    refer it to perform operations.
  - Alternatively, you can add the same methods in derived class that can accept
    mixed types and proxy to parent's method by passing the embedded value.
*/
package main

import "fmt"

// A CommonObject provides common fields for any object type
type CommonObject struct {
	Id   int
	Name string
}

// A Folder object that inherits CommonObject
// - Specify the name of a struct inside another struct for composition
type Folder struct {
	Level        int
	CommonObject // composition -- Folder inheriting all properties of CommonObject
}

// A File object inherting Folder object (multi level inhertiance)
type File struct {
	Folder
	Size int
}

// Print display the information of CommonObject
func (c *CommonObject) Print() {
	fmt.Println("\tCommon: Id =", c.Id, ", Name =", c.Name)
}

// Print display the information of Folder object
// This also overrides the Print method in CommonObject
// But, it can invoke inner struct's Print with the full name
func (f *Folder) Print() {
	f.CommonObject.Print()
	fmt.Println("\tFolder: Level =", f.Level)
}

// DemoStructEmbed demonstrates the struct composition in Go
func DemoStructEmbed() {
	fmt.Println("A struct can be embedded to another for inheritance")

	// Initializing a folder
	// Literals must use the full form to initialize embedded structs
	var folder Folder = Folder{Level: 1, CommonObject: CommonObject{Id: 1, Name: "folder"}}
	// Embedding provides shortcut access "folder.Id" as well as full "folder.CommonObject.Id"
	folder.Print()

	// Initialing file
	var file1 File = File{Size: 100, Folder: folder}
	// Attributes can be accessed with shorthand names
	fmt.Println("\tInherited attributes can be accessed directly:", file1.Id, file1.Name)
	// Attributes can also be accessed like 'file1.CommonObject.Name'
	// Not the best practice, but useful in case if the original and embedded type has the same attribute.

	// Literals must provide the full path
	var file2 File = File{Size: 200, Folder: Folder{Level: 2, CommonObject: CommonObject{Id: 2, Name: "folder2"}}}
	fmt.Println("\tUnfortunately no shorthand available for literals:", file2.Id, file2.Name)
	file2.Print() // invokes Folder.Print since File does not have Print()
}
