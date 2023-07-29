/*
Demontrates writing generic functions in Go to support multiple types
*/
package main

import "fmt"

// Number is a Go type constraint and is a union of int32 and float32 types.
// Type constraints are required while declaring Go generic's type parameters
// We can simply write "Number" wherever we want to write "int32 | float32"
type Number interface {
	int32 | float32
}

// SumNumbers produces the sum of given numbers
// [K comparable, V Number]: called Go type parameters
//   - comparable is Go's builtin interface to support all types that can use
//     operator == and !=
//   - Number is our own interface to support type constraint int32 | float32
//   - Each Type parameter is specified with a type constraint
//   - Type constraint allows set of permissible types like "int | float32"
//   - But it stands for a single type at compile time based on the calling
//     function.
//   - Function will be called with both type parameters and ordinary function
//     arguments.
//   - However, the type parameters can be omited in some cases where the
//     compiler can determine automatically
//   - Note that the multiple types must support all the operations performed
//     in the function.
//
// (numbers map[K]V): ordinary function arguments using the defined types
func SumNumbers[K comparable, V Number](numbers map[K]V) V {
	var sum V
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	ints := map[string]int32{
		"a": 1, "b": 2, "c": 3,
	}
	// Calling function can specify both type and ordinary arguments
	sumInts := SumNumbers[string, int32](ints)
	fmt.Println(sumInts)

	floats := map[string]float32{
		"x": 1.1, "y": 2.2, "z": 3.3,
	}
	// Calling function can omit type arguments since the compiler
	// can determine from the argument "floats map[string]float32"
	sumFloats := SumNumbers(floats)
	fmt.Println(sumFloats)

	// Type arguments can not be omitted always. For instance, when calling a
	// function that takes no ordinary arguments. Example:
	// func ReadNumber[T Number]() T
	// must be called with
	// intNum := ReadNumber[int32]()
}
