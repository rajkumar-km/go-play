/*
comma inserts a comma after every 3 digits
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("1234567283463629342"))
}

// comma inserts a comma after every 3 digits
// Example: 1234567 becomes 1,234,567
func comma(s string) string {
	var b bytes.Buffer

	// Write the initial section which can vary length
	// Example: 1 in (1,234,567)
	i := len(s) % 3
	if i == 0 {
		i = 3 // Start from 3rd position if the digits are multiples of 3
	}
	b.WriteString(s[:i])

	// Write the remaining sections which are always 3 digits
	// Example: 234 and 567 in (1,234,567)
	for ; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i:i+3])
	}
	return b.String()
}
