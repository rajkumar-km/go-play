/*
Package main demonstrates the control flow statements in Go
*/
package main

import "fmt"

// main demonstrates the conditional and looping statements in Go
// 1. "if" statement
// 2. "switch" statement
// 3. "for" statement
func main() {
	// If statement
	// Condition can be with or without parenthesis
	// A short variable declaration can be used in non parenthesis format
	if x := 25; x%5 == 0 {
		fmt.Printf("%d is the multiple of 5", x)
	} else if x%2 == 0 || x%9 == 0 {
		fmt.Printf("%d is the multiple of 2 or 9", x)
	} else {
		fmt.Printf("%d is not the multiple of 5 or 2 or 9", x)
	}

	// Switch statement
	// A short declaration is allowed in switch as well
	switch dayOfweek := 5; dayOfweek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid")
	}

	// Switch with no expression is like if-else-if
	var dayOfweek = 7
	switch {
	case dayOfweek > 0 && dayOfweek < 6:
		fmt.Println("Weekday")
	case dayOfweek < 8:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid")
	}

	// For loop - the only loop in Go
	// Usage1: for initialization ; condition ; increment {}
	// Usage2: for ; condition ; increment {}
	// Usage3a: for ; condition ; {}
	// Usage3b: for condition {}
	// Usage4: for {}
	// Usage5: for i,val := range(values) {}
	for i := 0; i < 10; i++ {
		// Break/Continue
		// break - break the loop
		// continue - proceed to next iteration

		if i == 1 {
			continue
		} else if i == 2 {
			break
		}
	}

	// For loop with range
	// Useful to iterate characters in string, arrays, slices, channels and maps
	// The blank identifier _ can be used in place of i if the index is not required
	for i, c := range "Hello" {
		fmt.Printf("%d: %c\n", i, c)
	}
}
