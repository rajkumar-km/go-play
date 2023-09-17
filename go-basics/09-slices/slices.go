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

func main() {
	createSlices()
	compareSlices()
	sliceFunctions()
	iterateSlices()
	inPlaceSliceOperations()
}

// main demonstrates the use of slices in Go
// 1. Creating slices from slice literals
// 2. Creating slices from array
// 3. Creating slices from other slice
// 4. Creating slices using make function
func createSlices() {

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
	// It allocates an underlying array with size equal to the given capacity,
	// and returns a slice that refers to that array. The extra capacity is reserved for
	// future extension although it is not accessible now
	fmt.Println(`Creating a slice using make`)
	var s []int
	fmt.Printf("\tSlices are by default initialized to nil. Either use make(), or a slice literal to initialize\n")
	fmt.Printf("\tvar s[]int, s == nil: %v\n", s == nil)
	if s == nil { // should be true
		s = make([]int, 5, 10)
	}
	fmt.Printf("\ts := make([]int, 5, 10)\n")
	fmt.Printf("\ts = %v, len=%d, cap=%d\n", s, len(s), cap(s))
}

// compareSlices demonstrates slice comparisions
// Unlike arrays, slices can not be compared because,
// 1. Elements of a slices are indirect. A slice can contain itself as an element.
// 2. Deep comparision is difficult since a slice holds references for map, pointers and channels.
//
// Even arrays can be compared with == operator only if it holds comparable elements.
// The == operator can not be applied on arrays containing slices, maps. Ex: [3][]string
func compareSlices() {
	fmt.Println(`Comparing slices`)

	// Slices can only be compared to nil
	var s1 = []string{"one", "two"}
	var s2 = []string{"one", "two"}
	// s1 == s2 // compile error: slices can only be compared to nil

	// Even arrays containing reference to slices can not be compared
	// var a1 [2][]string = [2][]string{s1, s2}
	// var a2 [2][]string = [2][]string{s1, s2}
	// fmt.Println(a1 == a2) // compile error: [2][]string can not be compared

	// So, write a function to compare your own slices
	fmt.Println("\tslices or even arrays containing slices can not be compared with ==")
	fmt.Println("\twrite own function compareSlice(s1, s2) =", compareSlice(s1, s2))

	// Always use len(s) to check for empty slices
	// Because a slice is not nil if initialized with empty literal []string{}
	var x []string
	var y = []string{}
	fmt.Printf("\tvar x []string\n\tvar y = []string{}\n")
	fmt.Printf("\tx == nil => %v\n\ty == nil => %v\n", x == nil, y == nil)
	fmt.Printf("\tlen(x) == 0 => %v\n\tlen(y) == 0 => %v\n", len(x) == 0, len(y) == 0)
}

// compareSlice compares two slices containing strings
func compareSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// sliceFunctions demonstrates using slice functions copy(), and append()
func sliceFunctions() {
	// Slice Functions
	// 1. The copy() function: copying a slice
	// The number of elements copied will be the minimum of len(src) and len(dst).
	// Usage: copy(dst, src []T) int
	// The arguments order is in dst = src assignment style
	src := []string{"one", "two", "three"}
	dst := make([]string, 2)
	numElementsCopied := copy(dst, src)
	fmt.Println(`Copying a slice`)
	fmt.Printf("\tcopy(dst(cap=2), src%v) = %v, nElementsCopied=%d\n", src, dst, numElementsCopied)

	// 2. The append() function: appending to a slice
	// Usage: append(s []T, x ...T) []T
	// - Accepts a slice and variable number of arguments to append
	// - Append to the same underlying array if the capacity is sufficient
	// - Otherwise creates a new base array, copy the old elements, append the new elements
	// - Generally append() allocates double the capacity to prevent from frequent allocations
	// - But we can not determine when it would allocate/create new array, so it is always
	//   safe to use the return value. s = append(s, ...)
	srcApp := append(src, "four", "five") // results in creating a new array since it extends the capacity of src
	fmt.Println(`Appending to a slice using append(s []T, x...T) []T`)
	fmt.Printf("\tsrc=%v, len=%d, cap=%d\n", src, len(src), cap(src))
	fmt.Printf("\tsrcApp=%v, len=%d, cap=%d\n", srcApp, len(srcApp), cap(srcApp))

	srcWithCapacity := make([]string, 3, 10)
	copy(srcWithCapacity, []string{"one", "two", "three"})
	srcWithCapacityApp := append(srcWithCapacity, "four", "five")
	srcWithCapacityApp[0] = "ONE" // Original array is affected since it has enough capacity
	fmt.Printf("\tsrcWithCapacity=%v, len=%d, cap=%d\n", srcWithCapacity, len(srcWithCapacity), cap(srcWithCapacity))
	fmt.Printf("\tsrcWithCapacityApp=%v, len=%d, cap=%d\n", srcWithCapacityApp, len(srcWithCapacityApp), cap(srcWithCapacityApp))

	// Appending to a nil slice also creates a new base array

	// Appending one slice to another can be done with ... operator
	// which expands a slice to list of arguments
	var myS1 = []string{"A", "B"}
	var myS2 = []string{"C", "D"}
	var myS3 = append(myS1, myS2...)
	fmt.Printf("\tA slice can be expanded as list of arguments using '...' operator\n")
	fmt.Printf("\tmyS3 := append(myS1, myS2...) appends slice2 with slice1: %v\n", myS3)
}

// iterateSlices shows how to loop over the slices
func iterateSlices() {
	// Slice of slices
	// Slices can be of any type. They can also contain other slices.
	mySofS := [][]string{
		{"A", "B"},
		{"C", "D"},
	}

	// Iterating over a slice using range operator
	fmt.Println(`Iterating slices`)
	for i, v := range mySofS {
		fmt.Printf("\t%d %v\n", i, v)
	}

	// Slice can also be iterated using indexes
	hello := []byte("hello")
	reverse(hello)
	fmt.Printf("\treverse(%q) = %q\n", "hello", hello)
}

// reverse reverses the given byte string
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// inPlaceSliceOperations shows how to perform slice operations in place
func inPlaceSliceOperations() {
	fmt.Println("In-Place slice operations")
	s := []string{"one", "", "three"}
	fmt.Printf("\ts = %#v\n", s)
	s = removeEmpties(s)
	fmt.Printf("\tAfter s = removeEmpties(s) = %#v\n", s)

	s2 := []string{"", "", "three", "", "five"}
	fmt.Printf("\ts2 = %#v\n", s2)
	s2 = removeEmptiesV2(s2)
	fmt.Printf("\tAfter s = removeEmptiesV2(s2) = %#v\n", s2)

	// Likewise stacks and queues can be implemented using slices.
}

// removeEmpties removes the empty strings from the given slice without making a copy
func removeEmpties(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

// removeEmptiesV2 removes the empty strings from the given slice without making a copy
// It takes the empty slice pointing to the same array and uses append()
func removeEmptiesV2(s []string) []string {
	res := s[:0] // empty slice pointing to the same base array
	for _, v := range s {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}
