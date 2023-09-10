package main

import (
	"bytes"
	"fmt"
)

// DemoBytes demonstrates the strings and byte slices
//   - The strings are immutable, so we can not modify it
//   - It can be converted to byte slices for modification and converted back
//   - However, every conversion takes a new copy which is memory inefficient
//   - So, a bytes.Buffer type can help to manipulate strings
func DemoBytes() {
	s := "Hello, World"
	b := []byte(s) // creates a new copy
	_ = string(b)  // converting back also creates a new copy

	// Package bytes has similar methods as strings to perform operations like
	// HasPrefix, HasSuffix, Contains, Join, Index, Fields, Count, and etc.,
	// But the arguments are []byte instead of string
	fmt.Printf("bytes.Contains = %v\n", bytes.Contains(b, []byte("World")))

	// A bytes.Buffer type can be used to efficiently perform operations on strings
	// Buffer starts from empty and can be added with byte, []byte, runes and strings.
	var buf bytes.Buffer
	buf.Write([]byte("Hello"))
	buf.WriteByte(',')
	buf.WriteString(" ")
	buf.WriteRune('த')
	buf.WriteRune('ம')
	buf.WriteRune('ி')
	buf.WriteRune('ழ')
	buf.WriteRune('்')
	fmt.Printf("bytes.Buffer = %s\n", buf.String())
}
