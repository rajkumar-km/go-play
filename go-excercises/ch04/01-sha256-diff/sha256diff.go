/*
sha256diff produces the 256 bits digest for two different stringsand counts the bits
that are different
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := []byte("X")
	y := []byte("x")
	sumX := sha256.Sum256(x)
	sumY := sha256.Sum256(y)
	fmt.Println(diffBits(sumX, sumY))
}

// PopCount shifts the bits one to count the set bits
func PopCount(x byte) int {
	c := 0
	for ; x != 0; c++ {
		x = x & (x - 1) // removes the rightmost set bit of x
	}
	return c
}

// diffBits returns the number of bits that are different between x and y
func diffBits(x [32]byte, y [32]byte) int {
	diff := 0
	for i := 0; i < len(x); i++ {
		// xor sets the bits that are different in x[i] and y[i]
		xor := x[i] ^ y[i]
		diff += PopCount(xor)
	}
	return diff
}
