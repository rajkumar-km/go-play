package main

import "fmt"

// init function is executed by default even before the main function.
// Some of the uses of init function
// - initializing network connections prior to execution
// - creating required files and directories
// - checking if the dependent resources available
// In fact, there can be more than one init functions for the same package
// and they are executed in the order defined.
// If multiple packages have init() then it is called as per the deep dependency
// loading. Finally, the init() on main package is called.
func init() {
	fmt.Println("init() function executed next")
}
func init() {
	fmt.Println("multiple init() functions are executed in order")
}

// DisplayName is a package level variable
var DisplayName string = initDisplayName()

// initDisplayName initializes the global variable DisplayName
// Global variables are initialized even before the init function
// So, the initDisplayName() is called even before init()
func initDisplayName() string {
	fmt.Println("global variable initialization first")
	return "Variables"
}
