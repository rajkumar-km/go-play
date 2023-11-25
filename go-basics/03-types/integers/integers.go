package main

import "fmt"

// DemoIntegers demonstrates the various use of int type in Go
//   - int literals - decimal, octal, and hexadecimal representations
//   - overflow
//
// Integer types:
//   - int8, int16, int32, int64, int (size depends on compiler)
//   - uint8, uint16, uint32, uint64, uint (size depends on compiler)
//   - byte - represents a piece of raw data. alias of int8 and can be used interchangably
//   - rune - represents unicode characters. alias of int32 and can be used interchangably
//   - uintptr - hold the pointers for low level programming (size not specified)
func DemoIntegers() {
	// int literals can be decimal, octal, or hexadecimal
	// i := 123 - decimal
	// j := 0345 - octal
	// k := 0xffffffff - hexadecimal
	fmt.Println("\nInteger literals")
	d := 123
	o := 0345
	h := 0xffffffff
	fmt.Printf("Dec  = %d\n", d)
	fmt.Printf("Oct  = %o %#[1]o\n", o)
	fmt.Printf("Hex  = %x %#[1]x %#[1]X\n", h)

	// Bytes and Runes are printed with %c, or with %q if quoting is desired:
	var ascii byte = 'a'
	unicode := 'ழ'
	newline := '\n'
	fmt.Printf("Byte = %d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("Rune = %d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("Rune = %d %[1]q\n", newline)       // "10 '\n'"

	// Unsigned overflow
	// The higher order bit is silently discarded if it exceeds the size
	// Note: uint8 is from 0 to 255
	fmt.Println("\nUnsigned overflow")
	var x uint8 = 0xFF
	fmt.Printf("x   = %08b (%d)\n", x, x)
	for i := 0; i < 3; i++ {
		x++
		fmt.Printf("x++ = %08b (%d)\n", x, x)
	}

	// Signed overflow
	// Signed numbers follows the 2's compliment form. The higher order bit is
	// allocated for sign (0 - positive, 1 - negative)
	// For signed type, the sign bit is affected. For example
	// Note: int8 is from -128 to 127
	// var y int8 = 127, and i+1 sets the sign bit so the value becomes -128
	fmt.Println("\nSigned overflow")
	var y int8 = 127
	fmt.Printf("y   = %08b (%[1]d)\n", y) // [1] refers first argument
	for i := 0; i < 3; i++ {
		y++
		fmt.Printf("y++ = %08b (%[1]d)\n", y)
	}

	// Arithmetic operators +, -, *, /, and % can be applied on integers.
	fmt.Println("\nOperators")

	// The behavior of / truncates the results towards zero if the operands are integers
	fmt.Printf("3/2   = %d\n", 3/2)
	fmt.Printf("3.0/2 = %f\n", 3/2.0)

	// The % operator can be applied only to integers.
	// - The behavior for negative numbers varies across programming languages
	// - In Go, the sign of the remainder is always same as the sign of the operand
	//   so -5%3 and -5%-3 are both -2.
	fmt.Printf("5%%3   = %d (%% always returns the sign of operand)\n", 5%3)
	fmt.Printf("5%%-3  = %d\n", 5%-3)
	fmt.Printf("-5%%3  = %d\n", -5%3)
	fmt.Printf("-5%%-3 = %d\n", -5%-3)

	// All the bitwise operations are supported only on integers
	// See bitwise.go for more details.
}
