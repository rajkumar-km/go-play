/*
interface demonstrates the use of interfaces in Go programming

An interface express generalization about the behaviour of other types.
It contains list of method declarations and provides abstraction.
Concrete types can implement all the methods to be compatiple with the
interface.
  - In Go, we don't have explicitly mention that this type implements the any particular
    interface. Simply satifying the interface methods are sufficient.
  - This is a special behavior in Go and its lets us define generic interfaces even for
    built in concrete types.

How does an interface type work with concrete values?
  - Under the hood, an interface value can be thought of as a tuple consisting
    of a value and a concrete type: [value, type]
  - Example: [{65, 94}, main.Rectangle]
  - If an interface method is called, the corresponding method on the concrete type is called

Assigning interface with pointer and value types:
  - Another point to remember is an interface can be assigned with both value and pointer type.
  - A copy is made and stored in interface when value type is stored. So any changes to the
    object is not reflected on the other side.
  - Storing a pointer type always reflects the changes other side.

Interface and methods with pointer receiver:
  - If the interface is stored with a value type object, then it can not invoke methods with
    pointer receiver.
  - Because, the value stored in the interface is not addressable, so it can not dereference
    it to pointer and invoke the method.
*/
package main

import "fmt"

// Shape abstracts the operations performed on different shapes
// - Defines the abstract behavior for similiar objects
// - For example, Circle, Rectangle, Square, and etc.,
type Shape interface {
	Area() float32
	Perimeter() float32
}

// A Rectangle is an implementation of Shape interface
// - Just need to implement all the methods of the interface. Thats all
// - No explicit specification is required in Go like Java/C++
type Rectangle struct {
	Length float32
	Width  float32
}

// Area returns the area of Rectangle
func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

// Perimeter returns the perimeter of Rectangle
func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}

// SumOfAreas accepts variable list of shapes and returns the total area
// An example of using interface types as arguments for reusability
func SumOfAreas(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += float64(s.Area())
	}
	return totalArea
}

// Play main the use of interfaces in Go
// - Create an implementation and assign to interface variable
// - Invoke the functions through interface
func main() {
	// An interface type can be assigned with any value that implements all of its methods
	var s Shape = &Rectangle{Length: 65, Width: 94}
	fmt.Printf("Type = %T, Value = %v, Area = %f\n", s, s, s.Area())
	var s2 Shape = &Rectangle{Length: 90, Width: 86}
	fmt.Printf("Type = %T, Value = %v, Area = %f\n", s2, s2, s2.Area())

	// We can use an interface type for variables
	// We can also have a Circle that implements Shape interface and stored
	// in the same variable.
	var shapes []Shape = []Shape{s, s2}
	fmt.Println("Sum of area =", SumOfAreas(shapes...))

	// This isn't the usual way. Just for demonstration
	// Assigning a value type to the interface makes a copy, so the original object remains
	var orgObj Rectangle
	orgObj.Length = 10.0
	var s3 Shape = orgObj     // makes copy
	copyObj := s3.(Rectangle) // retrieve the copy from interface variable
	copyObj.Length = 20.0     // updating copied obj does not reflect in original object
	fmt.Printf("Original value: %f, Interface copy: %f\n", orgObj.Length, copyObj.Length)
}
