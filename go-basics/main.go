/*
go-basics demonstrates the core Go programming concepts
1. Hello, World
2. Variables
3. Types
4. Constants
*/
package main

import (
	"fmt"

	"github.com/rajkumar-km/go-play/go-basics/constants"
	"github.com/rajkumar-km/go-play/go-basics/hello"
	"github.com/rajkumar-km/go-play/go-basics/types"
	"github.com/rajkumar-km/go-play/go-basics/variables"
)

// main function is executed by default when running the package 'main'
func main() {
	fmt.Println("Executing main function")

	fmt.Println("----- 1. Hello, world -----")
	hello.Play()

	fmt.Println("----- 2. Variables -----")
	variables.Play()

	fmt.Println("----- 3. Types -----")
	types.Play()

	fmt.Println("----- 4. Constants -----")
	constants.Play()
}
