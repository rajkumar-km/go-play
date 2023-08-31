/*
assignments demonstrates the different ways of assigning values to variables

	variable = expression
	var1, var2 = var2, var1

	- Note that all the expressions on the right side is evaluated first before
	  any assignment to variables. This helps to use the same variable both on
	  left and right hand side without any issues.
*/
package main

import "fmt"

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

	// Tuple assignment allows to multiple assignments at once
	// Note that all the expressions on the right side is evaluated first before
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