package main

import "fmt"

// DemoString demonstrates the string type in Go.
// - A string is a immutable sequence of bytes
// - It usually contains the human readable text UTF8, but can also contain any arbitrary data.
func DemoString() {
	s := "hello world"
	fmt.Printf("s = %q, len(s) = %d\n", s, len(s))
	fmt.Printf("s[0] = %q (s[i] - 0 <= i <= len(s))\n", s[0])
	fmt.Printf("s[3:5] = %q\n", s[3:5])
	fmt.Printf("s[:5] = %q\n", s[:5])
	fmt.Printf("s[6:] = %q\n", s[6:])
	fmt.Printf("s[:] = %q\n", s[:])
	fmt.Printf("s[0:len(s)] = %q\n", s[0:len(s)])

	// panic: index out of range
	// s[len(s)]

	// strings are immutable, and we can not change the characters
	// s[1] = 'a' // compile error: cannot assign to s[0]
	// So, it is safe to pass on the strings in same memory
	// A substring like s[6:] also shares the same memory

	fmt.Println("String concatenation produces new string:")
	// Use + operator to concat and make new strings
	fmt.Printf("\"bye \" + s[6:] = %s\n", "bye "+s[6:])

	// Since the + operator creates new strings, the old string is untouched
	// It will be garbage collected only if there are no reference to it
	t := s
	s += "!"
	fmt.Println(`t := s ; s += "!"`)
	fmt.Printf("t = %s (still has reference to old string)\n", t)
	fmt.Printf("s = %s (allocates new string)\n", s)

	// Within a raw string literal, no escape sequences are processed
	// - Raw string literals can span over multiple lines
	// - Carriage returns are deleted to align with multiple platforms
	// - Useful to hold regular expressions, HTML templates, JSON literals, and
	//   command usage help.
	fmt.Println(`Go string literals can include escapes:
	\a - alert or bell
	\b - backspace
	\f - form feed
	\t - tab
	\v - vertical tab
	\' - single quote (only within double quotes)
	\" - double quote (only within single quotes)
	\\ - backslash
	\xff - represents a hex value (2 letters)
	\377 - represents a oct value (3 digits, max 377)`)

	// Arbitrary bytes can also be part of string literal
	octInLiteral := "\377"
	hexInLiteral := "\xff"
	fmt.Println(octInLiteral, hexInLiteral)
}
