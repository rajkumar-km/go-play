/*
popcount counts the number of bits set (that is 1) in given number
This is called population count
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

// Let's define a table to maintain the population count of numbers
// from 0 to 256. This should cover a popcount for 8 bits. This table
// can be utilized to count any types (say an uint64 can be divided 8 bytes)
var pc [256]byte

// init initializes the population count from 0 to 256
// This will be executed before main and useful to initialize complex objects
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountV0 returns the population count (number of set bits) of x.
// PopCountV0 = less than 0.15 ns/op
func PopCountV0(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountV1 simplifies the PopCountV0 with a for range loop
// PopCountV1 = took 3.7 ns/op in benchmark
func PopCountV1(x uint64) int {
	var c int = 0
	for i := 0; i < 8; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return c
}

// main parses a number from the command line arguments and invoke
// PopCountV0 and PopCountV1
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <number>\n", os.Args[0])
		return
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid number: %s\n", os.Args[1])
		return
	}

	fmt.Println("PopCountV0 =", PopCountV0(n))
	fmt.Println("PopCountV1 =", PopCountV1(n))
}
