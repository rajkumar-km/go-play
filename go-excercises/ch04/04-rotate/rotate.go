/*
rotate performs rotations on slice
*/
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)

	// Rotate s left by two positions.
	leftRotate(s, 2)
	fmt.Println(s, "after leftRotate(s, 2)") // "[2 3 4 5 0 1]"

	// Rotate by 3 positions, but it can return a new slice
	s = leftRotateV2(s, 1)
	fmt.Println(s, "after leftRotateV2(s, 1)")
}

// leftRotate rotates the array n times on left
func leftRotate(s []int, n int) {
	t := make([]int, n)
	copy(t, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], t)
}

// leftRotateV2 rotates the array n times on left
// It can allocate new base array and returns the resulting slice
func leftRotateV2(s []int, n int) []int {
	return append(s[n:], s[:n]...)
}
