package main

import (
	"fmt"
	"strconv"
)

// DemoStrconv demonstrates using Go strconv package
// strconv package contains functions to convert between strings and numeric types
func DemoStrconv() {
	// Itoa or Sprintf
	i := 1234
	s := strconv.Itoa(i)
	fmt.Printf("strconv.Itoa(%d) = %q\n", i, s)

	_ = fmt.Sprintf("%d", i) // is an alternative
	fmt.Printf("fmt.Sprintf(\"%%d\", %d) = %q\n", i, s)

	// Atoi or Sscanf for int type
	i, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("strconv.Atoi(%q) = %d\n", s, i)
	}

	_, err = fmt.Sscanf(s, "%d", &i)
	if err == nil {
		fmt.Printf("fmt.Sscanf(%q, \"%%d\", &i) = %d\n", s, i)
	}

	// ParseInt is genertic and accepts base and bit size
	// For example, convert aabbccdd to base 16 and bit size 32
	x, err := strconv.ParseInt("aabb", 16, 32)
	if err == nil {
		fmt.Printf("parse hex = %x\n", x)
	}

	// FormatInt or (%b, %d, %x, %u)
	fmt.Println("octal conversion = ", strconv.FormatInt(10, 8))
}
