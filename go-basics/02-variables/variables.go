/*
variables demonstrates the different ways to declare and initialize
the Go variables.

A variable is a piece of storage containing a value
Usage:

	var name type = expression

	  - Either the type or expression can be omitted but not both

Examples:

	var counter int = 10
	var counter = 10
	var counter int

Short variable declaration (allowed only inside the functions)
Usage:

	name := expression

		- The type is automatically determined by the type of expression
		- Type inference sometimes may not satisfy the requirements. For example, float values
		  are always float64 by default. We still need "var" form to if we need float32 type.

Examples:

	s := "hello" // string
	i, f64 := 10, 20.3 // int, float64

Go initializes all the variables to its corresponding zero value by default:
  - Global variables are initialized before the main() starts
  - Function variables are initialized only when the function is invoked

Naming in Go:
The names of Go variables, constants, functions, types, statement labels, and packages:
  - Must start with a letter or underscore
  - Can be in any length with letters, digits, and underscores. But Go recommends shorter names.
  - Case sensitive: getName() and GetName() are two different functions
  - Go has 25 keywords that can not be used as names
  - Go programmers use the "camel case" style for names.
  - The acronyms are left as it is in the capital. For example, escapeHTML() is valid and not escapeHtml()

Go keywords that can not be used as names:
  - break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto
  - if, import, interface, map, package, range, return, select, struct, switch, type, var

Go predeclared names that can be reused:
The following are the predeclared Go names that aren't reserved. This can be used in the program
but wherever it really makes sense. Please be aware of the confusion caused by redeclaration.
  - Constants - true, false, iota, nil
  - Types - int, int8, int16, int32, int64, uint, uint16, uint32, uint64, float32, float64,
    complex64, complex128, byte, rune, bool, string, error
  - Functions - new, make, len, cap, copy, append
    close, delete, complex, real, imag, panic, recover
*/
package main

import "fmt"

// DemoVariables function demontrates the Go variables:
// - typed variables
// - type inference
// - short declarations
func DemoVariables() {

	// Typed variables (Go is statically typed language)
	var myInt int // auto initialized to 0
	var myFloat float32 = 3.14
	fmt.Println(myInt, myFloat)

	// Go automatically infers the type of the variable based on the value provided
	var myString = "auto type inference"
	fmt.Printf("%T\n", myString)

	// Multiple variables of the same type
	var x, y, z int = 10, 20, 30
	fmt.Println(x, y, z)

	// Multiple variables of different type must use auto type inference
	// Float values are float64 by default and we can not specify type in multiple variable
	// declarations in a single line
	var i, f64, bl = 40, 50.6, true
	fmt.Println(i, f64, bl)

	// A set of variables can be initialized by calling a function that returns multiple
	// return values
	// Example: var f, err = os.Open(name) // returns a os.File and an error

	// A var block can be used to simplify multiple variables
	var (
		i1     int
		s1, s2 string = "hello", "world"
	)
	fmt.Println(i1, s1, s2)

	// Short declaration with ":=" is allowed only inside the function
	// One can not easily understand the type from short declaration
	// So, this is deliberately not allowed for global variables to improve readability
	var1 := 20
	var2 := 22.5
	var3, var4 := "hello", "world"
	var5 := false
	fmt.Println(var1, var2, var3, var4, var5)

	// Another important point about short declaration:
	// - The short declaration may include a variable that is already declared
	// - But it must have at least one new variable
	f, errCode := "file", 1
	fmt.Println(f, errCode)

	// This woulld result in compile error: no new variables on the left side of :=
	// f, errCode := "file2", 2

	// This works since we have a new variable f2
	// errCode is just a assignment and not redeclaration
	f2, errCode := "file2", 2
	fmt.Println(f2, errCode)

	// But, please aware that the same code inside the nested block causes a reclaration
	// of errCode. Any changes inside the block is not reflected outside since they both
	// are different variables in different scope
	{
		// Tip: use assignment "=" operator instead of ":=" to reuse the errCode and
		// make the updates visible outside this block
		f3, errCode := "file3", 3
		fmt.Println(f3, errCode)
	}

	// expected "2", because a new variable is created inside the above block
	fmt.Println("expecting 2 =", errCode)
}
