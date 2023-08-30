/*
lifetime demonstrates the lifetime of variables in Go. This is different from scope.
  - Lifetime focuses on keeping the memory and time to garbage collect
  - Whereas, scope focuses on accessiblity of the variable

Lifetime of Variables:

  - The lifetime of package variable is the entire execution of the program
  - By contrast, local variables have dynamic lifetime:
  - They are created when the declaration statement is encounted and lives as long as it is
    reachable. It will be garbage collected when it becomes unreachable
  - Interestingly, we can reference a local variable inside a function and store it in a
    global variable. The local variable lives beyond the function all through the program
    execution
  - Stack or Heap memory allocation is here determined by the compiler and not by the
    way we declare variables.
  - Although the memory is managed by garbage collector, we still need to be aware of the
    lifetime of variables in order to write efficient programs. We must not reference a
    short lived objects in global variables because it would occupy the memory and garbage
    collector can not recliam it.
*/
package main

import "fmt"

// A package level variable lives all through the program execution
var CmdLevel *string // initialized to nil by default

// DemoLifetime demonstrates the lifetime of variables in Go
func DemoLifetime() {

	// A local variable can outlive the function if it is referenced somewhere
	// The value of variable "level" is not freed after this function. This is accessible
	// through CmdLevel. So compiler chooses to allocate this on Heap memory.
	level := "level1"
	CmdLevel = &level

	// A local variable even allocated with "new" is freed after this function.
	// So, compiler can choose stack memory for variable "local"
	var local *int = new(int)
	fmt.Println(*local)
}
