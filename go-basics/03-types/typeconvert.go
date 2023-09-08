/*
typeconvert demonstrates the type conversion, type assertions, and type switch in Go
*/
package main

import "fmt"

// DemoTypeConversion demonstrates Go type conversions
func DemoTypeConversion() {
	demoTypeConv()
}

// demoTypeConv demonstrates converting one type to another
// Perform type convertion in Go by T(v). Converts value v to type T.
// Unlike C, Go does not perform implicit type conversion for safety (like buffer overflow)
// Explicit type conversion is required for everything even for adding two numeric types.
// Conversion are allowed in numeric, string, and some slice types
// Conversion between incompatiple types are reported at compile time. So, in any case,
// a conversion never fails at run time
func demoTypeConv() {
	// conversion between numeric types
	var intVal = 5
	var floatVal = 3.14
	var sum = floatVal + float64(intVal)
	fmt.Printf("floatVal + float64(intVal) = %f\n", sum)

	// conversion is required even between int32 and init64
	var myInt32 int32 = 10
	var myInt64 int64 = 20
	var intSum int64 = int64(myInt32) + myInt64
	fmt.Printf("int64(myInt32) + myInt64 = %d\n", intSum)

	// values after the precision is truncated when converting floats to int
	fmt.Println(floatVal, "becomes", int64(floatVal), "after type conversion to int64")

	// string <=> []byte slice conversions
	// string to []byte allocates a copy of the data
	var s string = "hello"
	var b []byte = []byte(s)  // creates a copy
	var s2 string = string(b) // creates a copy
	b[0] = 'w'
	fmt.Println("string =", s2, "[]byte =", string(b))
}
