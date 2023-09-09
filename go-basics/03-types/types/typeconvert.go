/*
typeconvert demonstrates the type conversion, type assertions, and type switch in Go
*/
package main

import "fmt"

// DemoTypeConversion demonstrates Go type conversions
//   - Perform type convertion in Go by T(v). Converts value v to type T.
//   - Unlike C, Go does not perform implicit type conversion for safety (like buffer overflow)
//   - Explicit type conversion is required for everything even for adding two numeric types.
//   - Conversion are allowed in numeric, string, and some slice types
//   - Conversion between incompatiple types are reported at compile time. So, in any case,
//     a conversion never fails at run time
func DemoTypeConversion() {
	demoIntConv()
	demoStrConv()
}

// demoIntConv demonstrates the type conversions that involves integers
func demoIntConv() {

	// Most of the conversions do not change the value, except
	// - a conversion that narrows a big integer into a smaller one
	// - conversion from integer to floating-point or vice versa

	// conversion between numeric types
	fmt.Println("\nConversion between numeric types")
	var intVal = 5
	var floatVal = 3.14
	var sum = floatVal + float64(intVal)
	fmt.Printf("floatVal + float64(intVal) = %f\n", sum)

	// conversion is required even between int32 and init64
	var myInt32 int32 = 10
	var myInt64 int64 = 20
	var intSum int64 = int64(myInt32) + myInt64
	fmt.Printf("int64(myInt32) + myInt64 = %d\n", intSum)

	// Conversion between float and int truncates the precision towards 0
	f := 3.741
	i := int(f)
	fmt.Printf("int(%f) becomes %d\n", f, i)

	// Avoid if the operand is out of range for the target type.
	// The result is implementation dependent
	f = 1e100
	i = int(f) // result is implementaion dependent
	fmt.Printf("int(%g) becomes %d is implementation dependent\n", f, i)

	// Conversion between int16 to int8 changes values
	fmt.Println("\nConversion from int16 to int8 change values")
	for _, v := range []int16{-256, -257, -258, -345, -512} {
		fmt.Printf("int16 to int8: %016b (%[1]d) => (%d) %08[2]b\n", v, int8(v))
	}

	// Conversion between signed and unsigned
	fmt.Println("\nConversion between signed and unsigned types")
	v := uint16(0x10F0)
	fmt.Printf("v               = %032b %d\n", v, v)
	fmt.Printf("int8(v)         = %032b %d\n", int8(v), int8(v))
	fmt.Printf("uint32(int8(v)) = %032b %d\n", uint32(int8(v)), uint32(int8(v)))
}

// demoStrConv demonstrates converting string types
func demoStrConv() {

	// string <=> []byte slice conversions
	// string to []byte allocates a copy of the data
	var s string = "hello"
	var b []byte = []byte(s)  // creates a copy
	var s2 string = string(b) // creates a copy
	b[0] = 'w'
	fmt.Println("string =", s2, "[]byte =", string(b))
}
