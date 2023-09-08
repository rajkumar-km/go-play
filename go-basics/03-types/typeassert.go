package main

import "fmt"

// DemoTypeAssertion demonstrates Go type assertions
// - Assertain the actual type stored in an interface
// - Type switch statement
func DemoTypeAssertion() {
	demoTypeAssert()
	demoTypeSwitch()
}

// demoTypeAssert shows how to assertain the actual type stored in an interface.
// An interface{} variable is capable of storing data of any type. Go provides
// few ways to determine the actual type stored in the variable.
// Example to use when you are sure about the underlying type:
//
//	var data interface{} = 10
//	myInt := data.(int) // Note: If the data is not int, then it causes panic.
//
// Example to assert the type without causing panic
//
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
