/*
typeswitch demonstrates Go type switch to discriminate different types stored in interface

An empty interface{} is like a discriminated union that can hold any concrete type.
*/
package main

import (
	"fmt"
	"strings"
)

// demoTypeSwitch shows the Go type switch to handle multiple types
// The "switch data.(type)" can be used to check for multiple cases like "case int:"
func main() {
	var data interface{} = 12.34

	switch data.(type) {
	case int:
		fmt.Println("data is int")
	case float32:
		fmt.Println("data is float32")
	case float64:
		fmt.Println("data is float64")
	default:
		fmt.Println("unsupported type")
	}

	// Let's look at another example to convert any value to SQL type
	q := SQL("SELECT * from users where id = ? AND name = ?", 1, "Alice")
	fmt.Println(q)
}

// SQL prepares a SQL query by accepting arguments
func SQL(query string, args ...any) string {
	for _,arg := range args {
		var argStr string

		// Note the variable creation in type switch
		// x is a local variable of type "any" within the switch block
		// Type switch also creates a local variable "x" for every case that matches
		// the specific type for the case.
		switch x := arg.(type) {
		case bool:
			if x { // x is new local variable "bool" here
				argStr = "true"
			} else {
				argStr = "false"
			}
			// fallthrough // compile error: cannot fallthrough in type switch
		case int, uint:
			argStr = fmt.Sprintf("%d", x) // x is "any" here because it combines int and uint
		case string:
			argStr = fmt.Sprintf("%q", x) // x is string
		default:
			panic(fmt.Sprintf("unsupported type: %T, %v", x, x))
		}

		query = strings.Replace(query, "?", argStr, 1)
	}
	return query
}
