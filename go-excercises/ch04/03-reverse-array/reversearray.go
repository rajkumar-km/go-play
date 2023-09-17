/*
reversearray reverses the array [8]int in place
*/
package main

import "fmt"

func main() {
	a := [8]int{31, 30, 29, 28, 27, 26, 25, 24}
	fmt.Println("a =", a)
	reverse(&a)
	fmt.Println("reverse(a) =", a)
}

// reverse reverses an integer array of size 8
// Note that it can only accept *[8]int type and not possible to with different length
// So, it is convinient to use slices to work with all sizes
func reverse(a *[8]int) [8]int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return *a
}
