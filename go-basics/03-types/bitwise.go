package main

import (
	"fmt"
)

// DemoBitwise demonstrates the bitwise operations in Go.
//   - Bitwise operations works only on integers both positive and negative numbers.
//   - However, the operand for bitwise shift operators must be unsigned integers.
//   - Generally bitwise operations are performed on unsigned integers, but remember
//     that, the sign bit also involves in the shift when using signed numbers
//
// Bitwise operations:
//   - AND (x&y)
//   - OR (x|y)
//   - XOR (x^y)
//   - NOT (^x)
//   - LEFT SHIFT (x<<1)
//   - RIGHT SHIFT (x>>1)
//   - AND NOT (x&^y)
func DemoBitwise() {
	// An example of set operations using bitwise
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("   x = %08b\n", x)    // the set {1,5}
	fmt.Printf("   y = %08b\n", y)    // the set {1,2}
	fmt.Printf(" x&y = %08b\n", x&y)  // the intersection {1}
	fmt.Printf(" x|y = %08b\n", x|y)  // the union {1,2,5}
	fmt.Printf(" x^y = %08b\n", x^y)  // the symmetric difference {2,5}
	fmt.Printf("x&^y = %08b\n", x&^y) // the difference {5}. x-y
	fmt.Printf("y&^x = %08b\n", y&^x) // the difference {2}. y-x
	fmt.Printf("  ^y = %08b\n", ^y)   // NOT
	fmt.Printf("x<<1 = %08b\n", x<<1) // LEFT SHIFT
	fmt.Printf("y>>1 = %08b\n", y>>1) // RIGHT SHIFT
}
