/*
scope demonstrates the different kinds of variable scope n Go
  - A package variable or global variable is declared outside the function
  - Local variables are declared inside the function. Function arguments and
    return values are also considered local variables.
  - Variables declared inside blocks (such as if, switch, for) are considered
    separate scope.

Exported names:
  - Global variable starting with capital letter is exported and accessible even
    outside the current package using <currPackageName>.GlobalVar
  - Likewise any constants, and functions starting with capital letter is exported
    and it can be accessed from other packages.
*/
package main

import "fmt"

// internalCmdName is a package variable accessible across the current package
// This includes accessing from the other files that belongs to current package
var internalCmdName string = "variables"

// Global variable exported outside the package
var CmdName string = "scope"

// DemoScope demonstrates the variable scope in Go
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
