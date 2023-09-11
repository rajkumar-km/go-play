package main

import (
	"fmt"
	"time"
)

// DemoConstants demonstrates the the typed and untyped constants in Go
//   - Looks like variable declaration but use a const keyword
//   - Values can not be changed once initialized
//   - The underlying types are strings, numeric, and boolean types
//   - Constants are evaluated at compile time. So, the results of all arithmetic, comparision,
//     and logical operations applied to constants are are themselves constants and it will be
//     evaluated at compile time. Thus saves the run time.
func DemoConstants() {
	// All the literals in Go are constants without a name
	fmt.Println(1, 3.17, "Hello", 'C', false, 1+2.7i)

	// Constants with name
	const appName = "Hello"
	const myVersion float32 = 2.0
	// use all formats like variable declaration and initialization

	// UnTyped constants - constants without explicit type specification
	// - This includes: untyped bool, untyped integer, untyped rune, untyped float, untyped
	//   complex, and untyped string
	// - These untyped flavours determine the actual type for variables (say an untyped float
	//   becomes float64)
	// - These can be used for assignment with any compatiple types (say number 10 can be assigned
	//   to int64, float64, and etc.,)
	// - Only constants can be untyped. If untyped constants are assigned to variables, then the
	//   variable infers the specific type automatically
	const myUntypedNumber = 10
	const myUntypedStr = "Hello"
	var assign float32 = myUntypedNumber
	fmt.Println(assign)

	// Typed constants comparision
	const dur = 5 * time.Minute // Go infers as the type of time.Minute which is time.Duration
	const myTypedInt int = 10
	const myTypedStr string = "World"
	type RichString string                            // Typename alias
	const assign1 RichString = myUntypedStr           // untyped does not requires type conversion for compatiple types
	const assign2 RichString = RichString(myTypedStr) // typed string requires type conversion

	// Constants and Type inference
	// Go has the default type for untypes constants and it will be used when defining variables
	var myInt = 10         // int
	var myFloat = 10.3     // float64
	var myBool = false     // bool
	var myChar = 'X'       // rune
	var myStr = "Hello"    // string
	var myComplex = 3 + 5i // complex128
	fmt.Println(myInt, myFloat, myBool, myChar, myStr, myComplex)

	// Constant expressions.
	// Multple constants can also be combined like variables
	// Untyped constants can be mixed in expression as long as the types are compatiple
	const (
		mySum   = myUntypedNumber + 10.5
		myCharC = 'B' + 1
	)
	fmt.Println(mySum, myCharC)

	// Constants within a block
	// If type/value is not specified, it can infer the same from previous constants
	// It can not be omitted for the first constant
	const (
		X = 1.0 // untyped float
		Y       // untyped float of 1. Infers both type and value from previous constant X
		Z = 'a' // untyped rune
		A       // untyped rune of 'a'
	)

	// Overflow and underflow are reported as compile error during the conversion
	// Rounding is allowed for float/complex
	const c0 = float32(0xdeadbeef) // rounded up from 3735928559 to 3735928576
	// const c1 := int32(0xdeadbeef) // overflow
	// const c2 := float64(1e309)    // overflow
	// const c3 := uint(-1)          // underflow

}
