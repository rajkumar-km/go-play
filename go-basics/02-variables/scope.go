/*
scope demonstrates the different kinds of variable scope n Go
*/
package main

import "fmt"

// internalCmdName is a global variable accessible across the current package
// This includes accessing from the other files that belongs to current package
var internalCmdName string = "variables"

// Global variable starting with capital letter is exported and accessible even
// outside the current package using <currPackageName>.GlobalVar
// Likewise any constants, and functions starting with capital letter is exported
// and it can be accessed from other packages.
var CmdName string = "scope"

func DemoScope() {
	// Variable declared in function level is accessible only within the function
	var level string = "level1"
	fmt.Println(level) // level 1

	{
		// Adding blocks creates new scope
		// The same name "level" can be redeclared in another scope
		var level string = "level2"
		fmt.Println(level) // level 2
	}

	// A short variable declaration is allowed in if, switch, select, and for
	// statements. Also, it limits the scope of the variable within the block
	if level := "level3"; len(level) > 0 {
		fmt.Println("Short variable declaration is allowed in if statement")
		fmt.Println("Also, it limits the scope of the variable only inside the if statement")
		fmt.Println(level) // level 3
	}
}
