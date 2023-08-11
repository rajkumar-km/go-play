package main

import "fmt"

// cmdName is a global variable accessible across the current package
// This includes accessing from the other files that belongs to current package
var cmdName string = "variables"

// Global variable exposed even outside the current package
// It can be accessed from other packages using <currPackageName>.GlobalVar
var DisplayName string = initDisplayName()

// init function is executed by default before the main function
// In fact, there can be more than one init functions for the same package
// and they are executed in the order defined. Some of the uses of init function
// - initializing network connections prior to execution
// - creating required files and directories
// - checking if the dependent resources available
func init() {
	fmt.Println("init() function executed next")
}
func init() {
	fmt.Println("multiple init() functions are executed in order")
}

// initDisplayName initialized the global variable DisplayName
// Global variables are initialized even before the init function
func initDisplayName() string {
	fmt.Println("global variable initialization first")
	return "Variables"
}
