/*
init explains the order in which Go runs the program
  - Deep dependent packages are initialized first before the main package
  - For every dependent package, Global variables are initialized first, and
    init() functions are invoked in the order they are defined
  - Next, the same procedure is repeated for the main package
  - Finally, start the execution from the main() function from main package
  - Variables inside the functions are initialized only when they are invoked

Some of the uses of init function:
  - initializing network connections prior to execution
  - creating required files and directories
  - checking if the dependent resources available
*/
package main

import "fmt"

// init function is executed by default even before the main function.
func init() {
	fmt.Println("init() function executed next")
}

// There can be more than one init functions for the same package
// and they are executed in the order defined.
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
