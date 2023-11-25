/*
Package main demonstrates the various data types supported in Go

Go types fall into four categories:
 1. Basic types - boolean, numeric (integers, float, and complex), and strings
 2. Aggregate types - array and struct
 3. Reference types - pointers, slices, map, chan, func
 4. Interface types - interface
*/
package main

import "fmt"

// DemoTypes demonstrates the types of variables supported in Go
func DemoTypes() {
	// Integer types
	// int8, int16, int32, int64, int (32 bit on 32-bit arch, 64 bit on 64-bit arch)
	// uint8, uint16, uint32, uint64, uint (32 bit on 32-bit arch, 64 bit on 64-bit arch)
	var myInt int8 = 64
	var myOct = 034
	var myHex = 0xFF
	fmt.Printf("int = %d, oct = %#o, hex = %#x\n", myInt, myOct, myHex) // used # for printing with prefix 0-oct and 0x-hex

	// Character types - just uses the aliases of numeric types to distinguish the char types
	// byte - uint8 - represents ASCII character
	// rune - int32 - represents other UTF characters
	var myChar = 'A' // inferred as 'rune' which is the default for character types.
	var myByte byte = 'B'
	var myRune rune = '♥'
	fmt.Printf("myChar = %c (%d), myByte = %c (%d), myRune = %c (%d)\n", myChar, myChar, myByte, myByte, myRune, myRune)

	// Floating point types
	// float32 - 32-bit single precision floating point (1 - sign, 8 - exponent, 23 - mantissa)
	// float64 - 64-bit double precision floating point (1 - sign, 11 - exponent, 52 - mantissa)
	var myFloat = 3.173 // inferred as float64 by default
	var myFloat32 float32 = 2.4545
	var myFloat64 float64 = 5.144
	// use %g to automatically print exponents as necessary digits
	// use %f to set width / precision
	// use %e for scientific notation
	fmt.Printf("myFloat = %2.2f, myFloat32 = %e, myFloat64 = %g\n", myFloat, myFloat32, myFloat64)

	// Booleans
	var myBool = false
	var myBool2 bool = true
	fmt.Println(myBool, myBool2)

	// Strings
	var myStr1 = "Normal String\nCan have Escape characters\n" // Normal string
	var myStr2 = `Raw string can span multiple lines
	but can not use escape characters`
	fmt.Println(myStr1, myStr2)

	// Complex numbers
	// complex64 - both real and imaginary are float32
	// complex128 - both real and imaginary are float64
	var myComplex = 2 + 2i // Type inferred as complex128 by default
	var myComplex2 complex64 = 3 + 4i
	var myComplex3 = complex(myFloat, myFloat64) // always use same bits in real and imaginary part.
	fmt.Println(myComplex, myComplex2, myComplex3)
}
