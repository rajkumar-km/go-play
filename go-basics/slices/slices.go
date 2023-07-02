/*
Package slices demonstrate the use of slices in Go

Arrays are fixed in size which makes it not flexible. A slice is actually a
reference pointing to subset of an array. Slice has length and capacity so that
it can be variable in length. A slice can be extended up to its capacity and
not beyond
*/
package slices

import "fmt"

// Play demonstrates the use of slices in Go
// 1. Creating slices from slice literals
// 2. Creating slices from array and other slice
// 3. Creating slices using make function
// 4. Slice functions copy(), append()
func Play() {

	// 1. Creating a slice using a slice literal
	// Internally creates an array of size 5 and returns the slice reference to it
	mySlice := []string{"A", "B", "C"}
	fmt.Println(mySlice, "len=", len(mySlice), "cap=", cap(mySlice))

	// 2. Creating a slice from an array
	// slice := array[low:high] // slice points to low..high-1 elements
	// - default value of low=0 and high=len
	// - length is the number of elements in the slice
	// - capacity is the maximum size up to which the segment can grow - see example
	// Example:
	// Array = [0, 1, 2, 3, 4]
	// Slice =    [1, 2, 3]
	//            <--len-->
	//            <----cap--->
	myArray := [6]string{"India", "Russia", "China", "USA", "Canada", "Brazil"}
	mySlice1 := myArray[3:5] // North America
	fmt.Println(mySlice1, "len=", len(mySlice1), "cap=", cap(mySlice1))

	mySlice2 := myArray[:3] // Asia
	fmt.Println(mySlice2, "len=", len(mySlice2), "cap=", cap(mySlice2))

	mySlice3 := myArray[3:] // America
	fmt.Println(mySlice3, "len=", len(mySlice3), "cap=", cap(mySlice3))

	// Modifying a slice affects the original array since it is holding the reference
	mySlice3[0] = "U.S.A"
	fmt.Println(myArray, mySlice1) // notice the changes in original array and other slices

	mySlice4 := myArray[:] // All
	fmt.Println(mySlice4, "len=", len(mySlice4), "cap=", cap(mySlice4))

	// 3. Creating a slice from another slice
	// The new slice still referencing to the original array and uses its capacity
	// Multiple slices can be created from the same array and all holding same references
	mySliceOfSlice := mySlice1[1:2]
	fmt.Println(mySliceOfSlice, "len=", len(mySliceOfSlice), "cap=", cap(mySliceOfSlice))

	// 4. Creating a slice using the built-in make() function
	// The make function takes a type, a length, and an optional capacity.
	// It allocates an underlying array with size equal to the given capacity, and returns a slice that refers to that array.
	mySlice5 := make([]int, 5, 10)
	fmt.Println(mySlice5, "len=", len(mySlice5), "cap=", cap(mySlice5))

	// 5. Zero value of slices
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nil")
	}

	// Slice Functions
	// 1. The copy() function: copying a slice
	// The number of elements copied will be the minimum of len(src) and len(dst).
	// Usage: copy(dst, src []T) int
	src := []string{"one", "two", "three"}
	dst := make([]string, 2)
	numElementsCopied := copy(dst, src)
	fmt.Println("src=", src, "dst=", dst, "numElementsCopied=", numElementsCopied)

	// 2. The append() function: appending to a slice
	// Usage: append(s []T, x ...T) []T
	// - Accepts a slice and variable number of arguments to append
	// - Append to the same underlying array if the capacity is sufficient
	// - Otherwise creates a new base array, copy the old elements, append the new elements
	srcApp := append(src, "four", "five")
	srcApp[0] = "ONE" // Does not affect the original array since it creates a new array
	fmt.Printf("src=%v, len=%d, cap=%d\n", src, len(src), cap(src))
	fmt.Printf("srcApp=%v, len=%d, cap=%d\n", srcApp, len(srcApp), cap(srcApp))

	srcWithCapacity := make([]string, 3, 10)
	copy(srcWithCapacity, []string{"one", "two", "three"})
	srcWithCapacityApp := append(srcWithCapacity, "four", "five")
	srcWithCapacityApp[0] = "ONE" // Original array is affected since it has enough capacity
	fmt.Printf("srcWithCapacity=%v, len=%d, cap=%d\n", srcWithCapacity, len(srcWithCapacity), cap(srcWithCapacity))
	fmt.Printf("srcWithCapacityApp=%v, len=%d, cap=%d\n", srcWithCapacityApp, len(srcWithCapacityApp), cap(srcWithCapacityApp))

	// Appending to a nil slice also creates a new base array

	// Appending one slice to another can be done with ... operator
	// which expands a slice to list of arguments
	var myS1 = []string{"A", "B"}
	var myS2 = []string{"C", "D"}
	var myS3 = append(myS1, myS2...)
	fmt.Println(myS3)

	// Slice of slices
	// Slices can be of any type. They can also contain other slices.
	mySofS := [][]string{
		{"A", "B"},
		{"C", "D"},
	}
	fmt.Println(mySofS)

	// Iterating over a slice
	for i := 0; i < len(mySofS); i++ {
		fmt.Println(i, mySofS[i])
	}
	// Iterating over a slice using range operator
	for _, v := range mySofS {
		fmt.Println(v)
	}
}
