package main

import "fmt"

func main() {
	// Demo complex numbers
	fmt.Println("--- DemoComplex -------")
	DemoComplex()
}

// DemoComplex demonstrates the complex numbers in Go.
// - There two complex types supported:
//   - complex64 - uses float32 for real and imaginary part
//   - complex128 - uses float64
func DemoComplex() {
	// Functions complex(), real(), and imag()
	var x complex128 = complex(2, 3)
	fmt.Printf("x     = %v, real(x) = %g, imag(x) = %g\n", x, real(x), imag(x))

	// Complex literal
	y := 7 + 2i // or 2i + 7
	fmt.Printf("y     = %v, real(y) = %g, imag(y) = %g\n", y, real(y), imag(y))
	z := 4i // becomes 0+4i
	fmt.Printf("z     = 4i becomes = %v\n", z)

	// Operations on complex numbers
	fmt.Printf("x + y = %v\n", x+y)
	fmt.Printf("x * y = %v\n", x*y)

	// Equality == and !=
	// Two complex numbers are equal if both real and imaginary is same
	// math/cmplx provides functions to operate with complex numbers
}
