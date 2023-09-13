/*
anagram finds whether two strings are anagrams. That is they contains same letters in a different order
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: anagram <word-1> <word-2\n")
		os.Exit(1)
	}
	fmt.Printf("anagram(%q, %q) = %v\n", os.Args[1], os.Args[2], anagram(os.Args[1], os.Args[2]))
}

// anagram finds whether two strings are anagrams.
func anagram(s1, s2 string) bool {
	// Anagram strings must be in same length
	if len(s1) != len(s2) {
		return false
	}

	// If s1 and s2 are equal then it is anagrams
	if s1 == s2 {
		return true
	}

	// Record the character occurances of s1 in a map
	m := make(map[rune]int)
	for _,c := range(s1) {
		m[c]++
	}

	// Check the for same occurances in s2
	for _,c := range(s2) {
		n, ok := m[c]
		if !ok {
			fmt.Printf("s2 has '%c', but s1 does not have the match\n", c)
			return false
		}
		if n == 1 {
			delete(m, c)
		} else {
			m[c]--
		}
	}

	return len(m) == 0
}