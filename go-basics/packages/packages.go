/*
Package "packages" demontrate the use of packages in Go

DRY (Don't Repeat Yourself) principle is one of guidance of good quality software
Packages are the next steps (after functions) for code Reusability, Modularity, Maintainability
- Reduces naming conflicts. Same function name can be reused in other packages
- Speed up the compilation since it is required for only modified packages

Importing packages:
Go convention is have the directory name same as package name
- For example "rand" is the nested package name on import of the "math/rand"
- Anything (variable, type, func) that starts with the Capital letter is accessible after the import
- Anything that starts with the small letter are private to the package

The main package is special and is also called command. It is a executable where the programs starts.
*/
package packages

import (
	// Go built in packages
	"fmt"
	"math"
	"math/rand"

	// Creating own packages
	"github.com/rajkumar-km/go-play/go-basics/packages/mypkg"
	"github.com/rajkumar-km/go-play/go-basics/packages/mypkg/mynestedpkg"

	// Using third party modules
	// go get rsc.io/quote // Go will also add this new dependency to the go.mod file.
	// go mod tidy // Automatically add/remove package dependencies based on the usage
	"rsc.io/quote"
)

// Play demonstrates defining and using packages in Go
// 1. Using built in packages
// 2. Creating and using our own packages
// 3. Importing third party modules in our code
func Play() {
	// MaxInt64 is an exported name from math package
	fmt.Println("Max value of int64: ", int64(math.MaxInt64))

	// Float32 is an exported function from rand
	fmt.Println(rand.Float32())

	// Invoke MyFunc() from package mypkg
	mypkg.MyFunc()

	// Invoke MyNestedPkgFunc() from package mynestedpkg
	mynestedpkg.MyNestedPkgFunc()

	// Invoke third party function Go() from quote package
	fmt.Println(quote.Go())
}
