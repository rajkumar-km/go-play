/*
Package main demonstrates the use of maps in Go programming

A map is an unordered collection of key-value pairs.
Syntax:

	var m map[KeyType]ValueType
*/
package main

import "fmt"

// main demonostrates the use of maps in Go
// - Creating maps using var, make(), and map literals
// - Adding, accessing, and deleting items from map
// - Checking if a key exists in map
// - Iterating maps
func main() {

	// Create a map
	var myMap map[string]int
	if myMap == nil {
		fmt.Println("nil map")
	}

	// Create a map using the built-in make() function
	myMap = make(map[string]int)
	myMap["Sun"] = 0
	fmt.Println(myMap)

	// Create a map using a map literal
	emptyMap := map[int]string{}
	fmt.Println("emptyMap = ", emptyMap)

	weekDays := map[string]int{
		"Sun": 0,
		"Mon": 1,
		"Tue": 2,
	}
	fmt.Println(weekDays)

	// Adding items (key-value pairs) to a map
	// If you try to add a key that already exists in the map, then it will simply be overridden by the new value.
	weekDays["Wed"] = 3
	weekDays["Thu"] = 4
	weekDays["Fri"] = 5

	// Retrieving the value associated with a given key in a map
	// The default value is returned if the key does not exists (say 0 for int)
	fmt.Println(weekDays["Wed"], weekDays["invalid"])

	// Checking if a key exists in a map
	value, isExists := weekDays["invalid"]
	fmt.Println(value, isExists)

	// Deleting a key from a map
	delete(weekDays, "Sun")
	fmt.Println(weekDays)

	// Maps are reference types
	// Assigning a map to another holds only the reference

	// Iterating over a map
	for key, val := range weekDays {
		fmt.Println(key, val)
	}
}
