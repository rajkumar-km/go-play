/*
Package main demonstrates the use of interfaces in Go programming

An interface contains list of method declarations and provides abstraction.
Concrete types can be implement all the methods to be compatiple with the
interface.

How does an interface type work with concrete values?
  - Under the hood, an interface value can be thought of as a tuple consisting
    of a value and a concrete type: [value, type]
  - Example: [{65, 94}, main.Rectangle]
  - If an interface method is called, the corresponding method on the concrete type is called
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
func (r *Rectangle) Area() float32 {
	return r.Length * r.Width
}

// Perimeter returns the perimeter of Rectangle
func (r *Rectangle) Perimeter() float32 {
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

// MyDrawing represents a drawing and use an interface type as field
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

// Play main the use of interfaces in Go
// - Create an implementation and assign to interface variable
// - Invoke the functions through interface
func main() {
	// An interface type can be assigned with any value that implements all of its methods
	var s Shape = &Rectangle{Length: 65, Width: 94}
	fmt.Printf("Type = %T, Value = %v, Area = %f\n", s, s, s.Area())

	var s2 Shape = &Rectangle{Length: 90, Width: 86}
	fmt.Println("Sum of area =", SumOfAreas(s, s2))
}
