/*
Package arrays demonstrates the arrays in Go programming
Arrays are fixed in size and stored in continuous memory
Go array is a value type. So, assigning array to another variable makes a copy.
Pointers can be used to work with array references.
*/
package arrays

import "fmt"

// Play demontrates the use of arrays in Go programming
// - Array declaration and access by index
// - Initiazing array with array literal
// - Array iterations
// - Multidimensional arrays
func Play() {
	simpleArray()
	arrayLiteral()
	iterateArray()
	multidimensional()
}

// simpleArray declares an array and access by index
func simpleArray() {

	// Simple arrays
	var numbers [5]int
	var words [3]string
	var cnums [2]complex64
	// Arrays by default initialized to default values
	fmt.Println(numbers, words, cnums)

	// Access by index
	words[2] = "last"
	fmt.Println(words[2], words)
}

// arrayLiteral initializes array with Go array literals
func arrayLiteral() {
	// Initialize by array literal
	var a = [5]int{1, 2, 3, 4, 5}
	var a2 = [5]int{1, 2, 3} // index 4 and 5 will be auto initialized to zero
	var a3 = [...]int{1, 2, 3, 4}
	fmt.Println(a, a2, a3, len(a3))

	// Array’s length is part of its type
	// For example, [5]int and [4]int are two different data types
	// So can not be assigned to each other
	var x1 = [2]int{1, 2}
	var x2 = [3]int{1, 2, 3}
	fmt.Println(x1, x2)
	// x1 = x2 // Not possible since [2]int and [3]int are distinct types
}

// iterateArray iterates the array using a simple for loop or with range operator
func iterateArray() {
	// Arrays in Golang are value types (not the reference type)
	// So assigning a array to another one creates a new copy always
	// Pointers can be used for call by reference
	daysOfWeek := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	daysOfWeekFull := daysOfWeek // Copy object by default
	daysOfWeekFull[0] = "Sunday" // Changes affects only the new copy
	fmt.Println(daysOfWeek, daysOfWeekFull)

	// Iterating arrays
	for i := 0; i < len(daysOfWeek); i++ {
		fmt.Println(i, daysOfWeek[i])
	}
	// Iterate using range operator
	for i, day := range daysOfWeek { // use "_" instead of "i" if you don't want the index
		fmt.Println(i, day)
	}
}

// Multidimensional arrays in Golang
func multidimensional() {
	var dimensions = [4][2]float64{
		{86, 90},
		{65, 90},
		{94, 88},
		{62, 90},
	}
	fmt.Println(dimensions)
}
