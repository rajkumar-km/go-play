package main

import (
	"fmt"
	"math"
)

// DemoFloats demonstrates the floating-point numbers in Go.
// - Go follows the IEEE 754 standard implemented by all modern CPUs
// - There two float types supported:
//   - float32 - 6 decimal digits of precision (approximate)
//   - float64 - 15 digits of precision
//
// - math.MaxFloat32 is 3.4e38
// - math.MaxFloat64 is 1.8e308
// - The smallest positive values are near 1.4e-45 and 4.9e-324.
func DemoFloats() {
	// Floating point types
	// float32 - 32-bit single precision floating point (1 - sign, 8 - exponent, 23 - mantissa)
	// float64 - 64-bit double precision floating point (1 - sign, 11 - exponent, 52 - mantissa)
	var myFloat32 float32 = .4545  // Digits can be omitted before the decimal point
	var myFloat64 float64 = 53546. // precision can also be omitted
	var myFloat = 3.173            // inferred either float32 or float64 based on compiler/env
	fmt.Printf("%f (float32) = %b\n", myFloat32, math.Float32bits(myFloat32))
	fmt.Printf("%f (float64) = %b\n", myFloat64, math.Float64bits(myFloat64))
	fmt.Printf("%f (float)   = %b\n", myFloat, math.Float64bits(myFloat))

	// float64 can be preferred over float32, because a float32 is not really large
	// and can accumulate errors rapidly unless one is quite careful
	var f float32 = 1 << 24                                              // 16777216
	fmt.Println("float32(16777216) == float32(16777216)+1 : ", f == f+1) // "true"!

	// Very large and very small numbers are better written with scientific notation
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	// Printf formats for float
	// %g - automatically prints in most compact exponential form
	// %e - exponent. Use for scientific notation
	// %f - no exponent. use to set width / precision
	fmt.Printf("%%g = %g\n", Avogadro)
	fmt.Printf("%%e = %e\n", Avogadro)
	fmt.Printf("%%f = %f\n", Avogadro)

	// math package has functions for creating and detecting the special values defined by IEEE 754
	// But this can not be directly used for comparision.
	// For example, any comparision with NaN value is always false. Use IsNan() instead.
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	nan := math.NaN()
	fmt.Println("nan == nan:", nan == nan)
	fmt.Println("math.IsNaN(nan):", math.IsNaN(nan))

	// It is a good practice to return a boolean indicating if any float operations failed
	// Say a func floatCompute() (value float64, ok bool)
}
