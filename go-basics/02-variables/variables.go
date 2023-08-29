/*
Package main demonstrates the different ways to declare and initialize
the Go variables.

Go initializes all the variables to its corresponding zero value by default.

Also, explains the order in which Go runs the program
  - Deep dependent packages are initialized first before the main package
  - For every dependent package, Global variables are initialized first, and
    init() functions are invoked in the order they are defined
  - Next, the same procedure is repeated for the main package
  - Finally, start the execution from the main() function from main package
*/
package main

import "fmt"

// DemoVariables function demontrates the typed variables, type inference, and
// short declarations
func DemoVariables() {

	// Typed variables (Go is statically typed language)
	// Variables are initialized by default (0, 0.0, "", false)
	var myInt int
	var myFloat float32
	var myString string
	var myBool bool
	fmt.Println(myInt, myFloat, myString, myBool)

	// Multiple variables at once and initialization
	var (
		i1     int
		f1     float32 = 3.17
		s1, s2 string  = "hello", "world"
		b1     bool
	)
	fmt.Println(i1, f1, s1, s2, b1)

	// Type Inference (Go automatically infers the type of the variable based on the value provided)
	var myVar1 = 10
	var myVar2 = 12.5
	var myVar3 = "hello"
	var myVar4 = true
	fmt.Printf("%T, %T, %T, %T\n", myVar1, myVar2, myVar3, myVar4)

	// Short declaration with ":=" is allowed only inside the function
	// One can not easily understand the type from short declaration
	// So, this is deliberately not allowed for global variables to improve readability
	var1 := 20
	var2 := 22.5
	var3, var4 := "hello", "world"
	var5 := false
	fmt.Println(var1, var2, var3, var4, var5)
}
