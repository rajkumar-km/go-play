/*
assignments demonstrates the different ways of assigning values to variables

	variable = expression

Tuple assignments:

	var1, var2 = var2, var1
	f, err := os.Open("file.txt")

	- Note that all the expressions on the right side are evaluated first before
	  any assignment to variables. This helps to use the same variable both on
	  left and right hand side without any issues.
	- Avoid tuple assignments if the expressions are complex. It will difficult
	  to understand otherwise.

Operators that requires Tuple assignments:

	val, ok := m[key] // map lookup
	val, ok := x.(T)  // type assertion
	val, ok := <-ch   // channel receive

Use blank identifier to skip any value:

	_, ok := m[key]   // Just check if a key exists and don't want to retrieve value

Implicit assignments:
  - function arguments are assigned with passed value
  - return values are assigned to result values
  - literals assigns values to individual items
  - Example: consider a slice literal
    s := []string{"zero", "one", "two"}
  - This performs implicit assignments as follows:
    s[0] = "zero"
    s[1] = "one"
    s[2] = "two"

Assignability:
  - The following are the rules of variable assignments:
  - types must exactly match: for example a int32 and int64 are considered different
  - nil is allowed for interface and any reference type
  - Constants have more flexible assignability rules. It can be untyped and used in
    various constants. For example, an untyped int below can be used for float assignments,
    but anyway it is restricted only to the numeric assignments.
    const delta = 10
    var f float32 = 10
  - The "==" and "!=" works based on assignability rules. It does not work if the values
    does not match the same type
*/
package main

import (
	"fmt"
	"os"
)

func DemoAssignments() {
	// Simple assignment
	x := 1
	x = 10
	x = x + 1

	// Arithmetic and bitwise operators has the corresponding assignment operator
	x += 1 // is equal to x = x + 1
	x *= 2 // is equal to x = x * 2

	// Numeric variables support increment/decrement but only with postfix form
	x++
	x--

	// Tuple assignment allows multiple assignments at once
	// Note that all the expressions on the right side are evaluated first before
	// any assignment to variables. This helps to use the same variable both on
	// left and right hand side without any issues.
	a := 10
	b := 20
	a, b = b, a // swap without using a temporary variable

	// Look at the compact version gcd function in Go
	x = gcd(128, 63)
	fmt.Println("gcd(128,63) =", x)

	// Also, look at the nth fibonacci number example
	x = fib(5)
	fmt.Println("fib(5) =", x)

	// Tuple assignments for functions returning multiple values
	f, err := os.Open("unavail.txt")
	fmt.Println("file =", f, ", err =", err)

	// Blank identifier can be used to skip any value
	_, err = os.Stat("unavail.txt")
	fmt.Println("err =", err)

	// Tuple assignments for three operators
	// 1. map key lookup
	var m map[string]int = make(map[string]int)
	m["one"] = 1
	v, ok := m["one"]
	fmt.Println("map key lookup returned =", v, ok)

	// 2. type assertion
	var i interface{} = 30
	v, ok = i.(int)
	fmt.Println("type assertion for int =", v, ok)

	// 3. channel receive
	var ch chan int = make(chan int)
	// launch a goroutine that closes a channel without writing any data
	go func(c chan int) { close(c) }(ch)
	v, ok = <-ch
	fmt.Println("channel receive =", v, ok)

	// Implicit assignments happens function arguments, return values, and
	// for array, slice, struct, and function literals
	// For example, This performs implicit assignments as follows:
	//  s[0] = "zero"
	//  s[1] = "one"
	//  s[2] = "two"
	s := []string{"zero", "one", "two"}
	fmt.Println(s)
}

// gcd determines the greatest common divisor for x and y
// For example, the greatest common divisor of 5 and 15 is "5"
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println("gcd: x =", x, ", y =", y)
	}
	return x
}

// fib identifies the nth fibonacci number
func fib(n int) int {
	x, y := 0, 1
	for ; n > 0 ; n-- {
		x, y = y, x+y
	}
	return x
}