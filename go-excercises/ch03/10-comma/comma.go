/*
comma inserts a comma after every 3 digits
*/
package main

import "fmt"

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

// comma inserts a comma after every 3 digits
// Example: 1234567 becomes 1,234,567
func comma(s string) string {
	for i := len(s); i > 3; i -= 3 {
		s = s[:i-3] + "," + s[i-3:]
	}
	return s
}
