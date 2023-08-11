package main

import "fmt"

// DemoTypeConversion demonstrates Go type conversions
func DemoTypeConversion() {
	demoTypecast()
	demoTypeAssert()
	demoTypeSwitch()
}

// demoTypecast demonstrates converting one type to another
// Perform type convertion in Go by T(v). Converts value v to type T.
// Unlike C, Go does not perform implicit type conversion for safety (like buffer overflow)
// Explicit typecast is required for everything even for adding two numeric types.
func demoTypecast() {
	var intVal = 5
	var floatVal = 3.14	
	var sum = floatVal + float64(intVal)
	fmt.Printf("floatVal + float64(intVal) = %f\n", sum)

	var myInt32 int32 = 10
	var myInt64 int64 = 20
	var intSum int64 = int64(myInt32) + myInt64
	fmt.Printf("int64(myInt32) + myInt64 = %d\n", intSum)
}

// demoTypeAssert shows how to assertain the actual type stored in an interface.
// An interface{} variable is capable of storing data of any type. Go provides 
// few ways to determine the actual type stored in the variable.
// Example to use when you are sure about the underlying type:
//	var data interface{} = 10
//	myInt := data.(int) // Note: If the data is not int, then it causes panic.
// Example to assert the type without causing panic
//	myInt, ok := data.(int)
func demoTypeAssert() {
	var data interface{} = 10

	// Get the data as integer when you are sure about the underlying type.
	myInt := data.(int)
	fmt.Printf("data.(int) = %d\n", myInt)

	// panic: interface conversion: interface {} is int, not float32
	// myFloat := data.(float32)

	// Safe assertion without causing panic
	myFloat, ok := data.(float32)
	if ok {
		fmt.Printf("data.(float32) = %f\n", myFloat)
	} else {
		fmt.Printf("data is not float32\n")
	}	
}

// demoTypeSwitch shows the Go type switch to handle multiple types
// The "switch data.(type)" can be used to check for multiple cases like "case int:"
func demoTypeSwitch() {
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
}