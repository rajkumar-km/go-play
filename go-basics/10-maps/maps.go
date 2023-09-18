/*
maps demonstrates the use of maps in Go programming

A map is an unordered collection of key-value pairs.
Syntax:

	var m map[KeyType]ValueType

- KeyType must be comparable using == so that it can check if the key exists
- ValueType can be anything. It can also be another map
- Maps are reference types, so assigning a map to another holds only the reference
- A map typically points to a storage that can grow infinitely when the map grows.
*/
package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func main() {
	createMaps()
	accessMaps()
	iterateMaps()
	compareMaps()
	mapAsSets()
}

// createMaps demonstrates creating maps using var, make(), and map literals
func createMaps() {

	// Create a map
	fmt.Println("var m map[string]int")
	var m map[string]int
	if m == nil {
		fmt.Println("\tyields a nil map")
		fmt.Println("\tAll the operations are safe on nil map except assigning a value to a key")
		// m["key"] = 1 // compile error: assignment to entry in nil map
	}

	// Create a map using the built-in make() function
	fmt.Println("m := make(map[string]int)")
	m2 := make(map[string]int)
	fmt.Printf("\tinitializes the map data structure with minimum storage, len=%d\n", len(m2))
	fmt.Printf("\tA map can grow any size as long as the system has memory\n")
	fmt.Printf("\tThe size parameter in make() does not apply to map, because it is hard to explain capacity for a map\n")

	// Create a map using a map literals
	fmt.Println("m := map[int]string{}")
	m3 := map[int]string{}
	fmt.Printf("\tinitializes the map with map literal, len=%d\n", len(m3))

	fmt.Println(`weekDays := map[string]int{"Sun": 0, "Mon": 1, "Tue": 2}`)
	weekDays := map[string]int{
		"Sun": 0,
		"Mon": 1,
		"Tue": 2,
	}
	fmt.Printf("\tinitializes the map with map literal, len=%d\n", len(weekDays))
}

// accessMaps demonstrates adding, accessing, and deleting items from map
func accessMaps() {
	weekDays := map[string]int{"Sun": 0, "Mon": 1, "Tue": 2}

	// Adding items (key-value pairs) to a map
	// If you try to add a key that already exists in the map, then it will simply be overridden by the new value.
	fmt.Println("Assign values using: m[key] = value")
	weekDays["Wed"] = 3
	weekDays["Thu"] = 4
	weekDays["Fri"] = 5

	// Retrieving the value associated with a given key in a map
	// The default value is returned if the key does not exists (say 0 for int)
	fmt.Println("Access the values using: m[key]")
	fmt.Println("\tMap access returns zero value if the key is not found")
	fmt.Printf("\tweekDays[\"unknown\"] = %d\n", weekDays["unknown"])
	fmt.Printf("\tSo, it is valid to perform operatons like weekDays[\"unknown\"]++\n")
	fmt.Printf("\tBut, we can take the address of a map value like &m[key], because a map is free relocate its elements\n")

	// Checking if a key exists in a map
	fmt.Println("Check if the key exists using: value, ok := weekDays[\"unknown\"]")
	fmt.Println("\tThe second value ok returns true if the key exists")
	value, ok := weekDays["unknown"]
	fmt.Printf("\tvalue, ok := weekDays[\"unknown\"] = %d, %v\n", value, ok)

	// Deleting a key from a map
	fmt.Println("Delete map keys using: delete(m, key)")
	delete(weekDays, "Sun")
	fmt.Println("\tweekDays after deleting Sun = ", weekDays)
}

// iterateMaps demonstrates iterating maps
func iterateMaps() {
	persons := map[string]int{"Charlie": 1, "Alice": 2, "Bob": 0}

	// Iterating over a map
	fmt.Println("Iterate maps using range, but it is unordered")
	for key, val := range persons {
		fmt.Println("\t", key, val)
	}

	// Ordered iteration by sorting keys
	fmt.Println("You can gather keys, sort it, and iterate them to access map")
	keys := make([]string, 0, len(persons))
	for k := range persons {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("\t", k, persons[k])
	}
}

// compareMaps shows how to compare two maps
func compareMaps() {
	fmt.Printf(`Like slices, maps can not be compared with == operator
So, we need to write our own function to compare maps
	persons := map[string]int{"Charlie": 1, "Alice": 2, "Bob": 0}
	comparePersons(persons, persons) = `)

	persons := map[string]int{"Charlie": 1, "Alice": 2, "Bob": 0}
	fmt.Println(comparePersons(persons, persons))
}

// comparePersons compares two maps for equality
func comparePersons(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// mapAsSets shows how to use maps like sets since Go does not have separate set structure
func mapAsSets() {
	fmt.Println("Map can be used like sets: map[string]bool or map[string]struct{}")
	in := "hi this is for testing words and it removes the the duplicates words"
	fmt.Print("\tBefore: ", in, "\n")
	fmt.Print("\tAfter :")

	r := strings.NewReader(in)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	set := make(map[string]bool)
	for scanner.Scan() {
		word := scanner.Text()
		if !set[word] {
			set[word] = true
			fmt.Print(" ", word)
		}
	}
}
