/*
comma_floats inserts a comma after every 3 digits and works with floats as well
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345.67"))
	fmt.Println(comma("1234.567"))
	fmt.Println(comma("12.34567"))
	fmt.Println(comma("-123.4567"))
}

// comma inserts a comma after every 3 digits
// Example: 1234567 becomes 1,234,567
// This supports floating point numbers and an optional sign
func comma(s string) string {
	var b bytes.Buffer

	// Check if the sign '-' is available. If so, write the sign and skip
	if s[0] == '-' {
		b.WriteByte('-')
		s = s[1:]
	}

	// Check if it is floating-point. If so, mark the boundary with lastIdx.
	precision := strings.LastIndex(s, ".")
	lastIdx := precision
	if lastIdx == -1 {
		lastIdx = len(s)
	}

	// Write the initial section which can vary length
	// Example: 1 in (1,234,567)
	i := lastIdx % 3
	if i == 0 {
		i = 3 // Start from 3rd position if the digits are multiples of 3
	}
	b.WriteString(s[:i])

	// Write the remaining sections which are always 3 digits
	// Example: 234 and 567 in (1,234,567)
	for ; i < lastIdx; i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i:i+3])
	}

	// Finally write the precsion as it is
	if precision != -1 {
		b.WriteString(s[precision:])
	}

	return b.String()
}
