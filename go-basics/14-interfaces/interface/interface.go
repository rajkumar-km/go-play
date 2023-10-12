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

Naming the interface:
  - Go recommends to name them with -er suffix, like Reader, Writer, Closer, ReadWriter, and etc.,

Interface Satisfaction:
  - All the methods of the interface must be implemented in concrete type.
  - It causes compile error if any of the interface methods are not implemented
  - The rule works even when assigning one interface value to another interface variables.
  - For example,
    var r io.Reader = new(bytes.Buffer)
    var rwc io.ReadWriteCloser = os.Stdout
    r = rwc
    rwc = r // compile error: bytes.Buffer lacks Close() method
  - Reader interface can be assigned with a value from ReadWriteCloser interface, because it
    satisfies the Read() method.
  - The other way is not possible, because bytes.Buffer type does not have Close() method.
  - The concrete type can have more methods than the ones specified by the interface. Once
    the concrete type is assigned to interface variable, it wraps and allows to call only
    the methods defined by the interface.
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
// - We can also have a Circle that implements Shape interface and stored in the same variable.
type Rectangle struct {
	Length float32
	Width  float32
}

// In Go, we don't have to associate interface and concrete type
// But, sometime it is useful to document the association and assertain the same at compile time
var _ Shape = (*Rectangle)(nil) // or "var _ Shape = Rectangle{}" but that creates an object

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

	var shapes []Shape = []Shape{s, s2}
	fmt.Println("Sum of area =", SumOfAreas(shapes...))

	// This isn't the usual way. Just for demonstration
	// Assigning a value type to the interface makes a copy, so the original object remains
	var orgObj Rectangle
	orgObj.Length = 10.0
	var s3 Shape = orgObj     // makes copy. values stored in the interface are not addressable
	copyObj := s3.(Rectangle) // retrieve the copy from interface variable
	copyObj.Length = 20.0     // updating copied obj does not reflect in original object
	fmt.Printf("Original value: %f, Interface copy: %f\n", orgObj.Length, copyObj.Length)
}
