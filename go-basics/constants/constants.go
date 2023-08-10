/*
main demonstrates the constants in Go programming
*/
package main

import "fmt"

// DemoConstants demonstrates the the typed and untyped constants in Go
func DemoConstants() {
	// All the literals in Go are constants without a name
	fmt.Println(1, 3.17, "Hello", 'C', false, 1+2.7i)

	// Constants with name
	const appName = "Hello"
	const myVersion float32 = 2.0
	// use all formats like variable declaration and initialization

	// UnTyped constants - constants without explicit type specification
	// These can be used for assignment with any compatiple types (say number 10 can be assigned to int64, float64, and etc.,)
	const myUntypedNumber = 10
	const myUntypedStr = "Hello"
	const assign float32 = myUntypedNumber
	fmt.Println(assign)

	// Typed constants comparision
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

	var divided = 10 / 4
	fmt.Printf("10/4 = %v (%T)\n", divided, divided)
	// Either use float number or explicit cast to float type to get the correct result
	var divided2 = 10.0 / 4
	fmt.Printf("10.0/4 = %v (%T)\n", divided2, divided2)

	var result = 4.5 + 5/2 // returns 6.5
	fmt.Printf("4.5 + 5/2 = %v (%T)\n", result, result)
	// Either use float number or explicit cast to float type to get the correct result
	var result2 = 4.5 + float64(5)/2 // returns 7.0
	fmt.Printf("4.5 + float64(5)/2 = %v (%T)\n", result2, result2)
}
