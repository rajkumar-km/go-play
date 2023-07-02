/*
go-basics demonstrates the core Go programming concepts
1. Hello, World
2. Variables
3. Types
4. Constants
5. Control Flow
6. Functions
7. Packages
8. Arrays
9. Slices
10. Maps
11. Pointers
12. Structures
13. Methods
14. Interfaces
15. Composition
16. Concurrency
*/
package main

import (
	"fmt"

	"github.com/rajkumar-km/go-play/go-basics/arrays"
	"github.com/rajkumar-km/go-play/go-basics/constants"
	"github.com/rajkumar-km/go-play/go-basics/controlflow"
	"github.com/rajkumar-km/go-play/go-basics/functions"
	"github.com/rajkumar-km/go-play/go-basics/hello"
	"github.com/rajkumar-km/go-play/go-basics/packages"
	"github.com/rajkumar-km/go-play/go-basics/types"
	"github.com/rajkumar-km/go-play/go-basics/variables"
)

// main function is executed by default when running the package 'main'
func main() {
	fmt.Println("Executing main function")

	fmt.Println("\n----- 1. HELLO WORLD -----")
	hello.Play()

	fmt.Println("\n----- 2. VARIABLES -----")
	variables.Play()

	fmt.Println("\n----- 3. TYPES -----")
	types.Play()

	fmt.Println("\n----- 4. CONSTANTS -----")
	constants.Play()

	fmt.Println("\n----- 5. CONTROL FLOW -----")
	controlflow.Play()

	fmt.Println("\n----- 6. FUNCTIONS -----")
	functions.Play()

	fmt.Println("\n----- 7. PACKAGES -----")
	packages.Play()

	fmt.Println("\n----- 8. ARRAYS -----")
	arrays.Play()
}
