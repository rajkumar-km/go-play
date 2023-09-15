/*
Package main demonstrate the use of slices in Go

Arrays are fixed in size which makes it not flexible. A slice is actually a
reference pointing to subset of an array. Slice has a pointer to base array, length
and capacity so that it can be variable in length. A slice can be extended up to its capacity and
not beyond.

Assigning a slice to another variable is like an alias. Or passing a slice only passes the reference.
The copy() function can be used to copy the contents of a slice

However, the append() function can be used to append elements beyond its
capacity. As a result, a new underlying array is created with double the
capacity and resultant slice is returned.

A slice can be accessed with a specific index s[i] or ranges s[m:n].
  - Like most programmmings, all indexing in Go uses half open intervals
    because it simplifies the logic
  - For example, s[m:n] is 0 <= m <= n <= len(slice) and contains n-m elements.
*/
package main

import "fmt"

// main demonstrates the use of slices in Go
// 1. Creating slices from slice literals
// 2. Creating slices from array and other slice
// 3. Creating slices using make function
// 4. Slice functions copy(), append()
func main() {

	// 1. Creating a slice using a slice literal
	// Internally creates an array of size 5 and returns the slice reference to it
	mySlice := []string{"A", "B", "C"}
	fmt.Println(`Creating new slice from literal`)
	fmt.Printf("\ts = %#v, len=%d, cap=%d\n", mySlice, len(mySlice), cap(mySlice))

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
	fmt.Println(`Creating slices from array`)
	myArray := [6]string{"India", "Russia", "China", "USA", "Canada", "Brazil"}
	fmt.Printf("\ta      = %#v, len=%d, cap=%d\n", myArray, len(myArray), cap(myArray))

	asia := myArray[:3] // Asia
	fmt.Printf("\tasia         = a[:3]  = %v, len=%d, cap=%d\n", asia, len(asia), cap(asia))

	america := myArray[3:] // America
	fmt.Printf("\tamerica      = a[3:]  = %v, len=%d, cap=%d\n", america, len(america), cap(america))

	all := myArray[:] // All
	fmt.Printf("\tall          = a[:]   = %v, len=%d, cap=%d\n", all, len(all), cap(all))

	// Modifying a slice affects the original array since it is holding the reference
	america[0] = "U.S.A"
	fmt.Println("Modifying a slice affects the base array so it reflects in other slices that uses the same base array")
	fmt.Printf("\tamerica[0]   = \"U.S.A\"\n")
	fmt.Printf("\ta            = %v\n", myArray)
	fmt.Printf("\tall          = %v\n", all) // notice the changes in original array and other slices
	fmt.Printf("\tamerica      = %v\n", america)

	// 3. Creating a slice from another slice
	// The new slice still referencing to the original array and uses its capacity
	// Multiple slices can be created from the same array and all holding same references
	northAmerica := america[:2] // North America
	fmt.Println(`New slices can be created from existing slices`)
	fmt.Printf("\tnorthAmerica = america[:2] = %v, len=%d, cap=%d\n", northAmerica, len(northAmerica), cap(northAmerica))

	allAfterAsia := asia[:6]
	fmt.Println("\t// It can extend the current slice, but within the capacity")
	fmt.Printf("\tasia         = a[:3]  = %v, len=%d, cap=%d\n", asia, len(asia), cap(asia))
	fmt.Printf("\tallAfterAsia = asia[:6] = %v, len=%d, cap=%d\n", allAfterAsia, len(allAfterAsia), cap(allAfterAsia))

	// 4. Creating a slice using the built-in make() function
	// The make function takes a type, a length, and an optional capacity.
	// It allocates an underlying array with size equal to the given capacity, and returns a slice that refers to that array.
	fmt.Println(`Creating a slice using make`)
	var s []int
	fmt.Printf("\tSlices are by default initialized to nil. Either use make(), or a slice literal to initialize\n")
	fmt.Printf("\tvar s[]int, s == nil: %v\n", s == nil)
	if s == nil { // should be true
		s = make([]int, 5, 10)
	}
	fmt.Printf("\ts := make([]int, 5, 10)\n")
	fmt.Printf("\ts = %v, len=%d, cap=%d\n", s, len(s), cap(s))

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
	
	// Iterating over a slice using range operator
	for _, v := range mySofS {
		fmt.Println(v)
	}

	// Slice can also be iterated using indexes
	hello := []byte("hello")
	reverse(hello)
	fmt.Printf("reverse(%q) = %q\n", "hello", hello)
}

// reverse reverses the given byte string
func reverse(s []byte) {
	for i,j:=0,len(s)-1; i < j; i,j=i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}