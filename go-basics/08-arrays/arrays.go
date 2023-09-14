/*
arrays demonstrates the arrays in Go programming
  - Array declaration and access by index
  - Initiazing array with array literal
  - Comparision
  - Array iterations
  - Multidimensional arrays

Arrays are fixed in size and stored in continuous memory
  - Go array is a value type. So, assigning array to another variable makes a copy.
  - Pointers can be used to work with array references.

Arrays are used rarely in Go for the fixed length operations such as sha256.
  - Slices are more convinient and often used. Anyway it is important to understand
    arrays since they are the foundation for slices.
*/
package main

import "fmt"

func main() {
	fmt.Println("\n--- Array Declaration -------")
	arrayDeclaration()

	fmt.Println("\n--- Array Literal -------")
	arrayLiteral()

	fmt.Println("\n--- Array Comparision -------")
	compareArray()

	fmt.Println("\n--- Iterating Array -------")
	iterateArray()

	fmt.Println("\n--- Multidimensional Array -------")
	multidimensional()

	fmt.Println("\n--- Passing Array as Arguments -------")
	arrayArguments()
}

// arrayDeclaration declares an array and access by index
func arrayDeclaration() {

	// Simple arrays
	var numbers [5]int
	var words [3]string
	var cnums [2]complex64
	// Arrays by default initialized to default values
	fmt.Println(numbers, words, cnums)

	// Access by index
	words[2] = "last"
	fmt.Println(words[2], words)

	// Find length of the array using "len" function
	fmt.Println(len(numbers))
}

// arrayLiteral initializes array with Go array literals
func arrayLiteral() {
	// Specify length
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Printf(`[5]int{1, 2, 3, 4, 5} => %#v`+"\n", a)

	var a2 = [5]int{1, 2, 3} // index 4 and 5 will be auto initialized to zero
	fmt.Printf(`[5]int{1, 2, 3}       => %#v`+"\n", a2)

	// var a3 = [2]int{1, 2, 3} // compile error: index 2 is out of bounds

	// Auto determine array length using [...]
	var a3 = [...]int{1, 2, 3, 4}
	fmt.Printf(`[...]int{1, 2, 3, 4}  => %#v, len=%d`+"\n", a3, len(a3))

	// Array literal with index-value pairs
	// Index can be specified in any order, and something can be ignored which
	// will be initiated to default value
	var a4 = [10]string{9: "nine", 5: "five"}
	fmt.Printf(`[10]string{9: "nine", 5: "five"} => %#v`+"\n", a4)

	// The [...] allocates the maximum index size in case index-value pairs are used
	var a5 = [...]byte{122: 'z'}
	fmt.Printf(`[...]byte{122: 'z'} => %#v`+"\n", a5)
}

// compareArray performs array comparisions
func compareArray() {
	// Compare the array with == operator. Returns true if both arrays are equal type
	// and has the same data.
	var a1 = [2]int{1, 2}
	var a2 = [2]int{1, 2}
	var a3 = [2]int{2, 1}
	fmt.Println(a1 == a2, a2 == a3) // true false

	// var a4 = [3]int{1, 2, 3}
	// if a3 == a4 // compile error: mismatched type [2]int and [3]int

	// Array’s length is part of its type
	// For example, [2]int and [3]int are two different data types
	// So can not be assigned or compared to each other
	// a3 = a4 // Not possible since [2]int and [3]int are distinct types
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

// arrayArguments demonstrates passing arrays as value vs reference
func arrayArguments() {
	x := [2]string{"Hello", "World"}
	fmt.Printf("x = %#v\n", x)
	passByValue(x)
	fmt.Printf("After passbyValue,     x = %#v\n", x)
	passByReference(&x)
	fmt.Printf("After passbyReference, x = %#v\n", x)

	// Although arrays can be passes with pointers, still it is not convinient for
	// passing variable length arrays. For example, *[2]string and *[3]string are
	// considered different types.
	// y := [3]string{"one", "two", "three"}
	// passByReference(&y) // compile error:
}

// passByValue accepts the array as argument and makes a copy
func passByValue(words [2]string) {
	// Go makes a copy of original array while passing it to function
	// So, the following statement does not affect the original array
	words[1] = "Go"
}

// passByReference accepts pointer to array so that it does not copy the actual array
func passByReference(words *[2]string) {
	// In Go, all the function arguments are pass by value
	// However, here it copies only the address of pointer variable. Still the
	// address points to the original array. So the following statement affects
	// the original array
	words[1] = "Go"
}
