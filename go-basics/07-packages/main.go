/*
packages demontrate the use of packages in Go

Package initialization:
  - Packages are initialized as per the import order. However, if they have
    dependencies, then the deep dependent packages are initialized first.
  - If a package has multiple files, to go tool sorts them by the name.
  - For every package, Go begins with initializing package variables in the order
    they are defined except that the dependencies are resolved first.
    var a = 10      // initialized first
    var b = c - 10  // initialized third because it has dependency of c
    var c = f()     // initialized second after calling function f
    func f() int { return a + 10 }
  - For some complex variables such as holding a table of values, an init
    function may be used.
  - Next, the same procedure is repeated for the main package
  - Finally, start the execution from the main() function from main package
  - Variables inside the functions are initialized only when they are invoked

The init() function:
  - An init() function a special function that is called automatically before main().
    It can not be called or referenced externally.
  - Multiple init() functions can be defined in the same package and they are invoked
    in the order they are defined
  - Example uses:
  - Useful to initialize complex data such building a map table which can not be a
    single liner.
  - Module developers can make use of this if necessary since they don't have main()
  - initializing network connections prior to execution
  - creating required files and directories
  - checking if the dependent resources available

The main package is special and is also called command. It is a executable where the programs starts.
*/
package main

import (
	// Go built in packages
	"fmt"
	"math"
	"math/rand" // a nested package under math
	"os"
	"strconv"

	// Using third party modules
	// go get rsc.io/quote // Go will also add this new dependency to the go.mod file.
	// go mod tidy // Automatically add/remove package dependencies based on the usage
	"rsc.io/quote"

	// Creating and importing own packages
	"github.com/rajkumar-km/go-play/go-basics/07-packages/bmi"

	// using alias name for packages to give a shorter name or
	// Avoid name conflicts if multiple packages presents with the same name
	bmiguide "github.com/rajkumar-km/go-play/go-basics/07-packages/bmi/guide"
)

// main demonstrates defining and using packages in Go
// 1. Using built in packages
// 2. Creating and using our own packages
// 3. Importing third party modules in our code
func main() {
	{
		// MaxInt64 is an exported name from math package
		fmt.Println("Max value of int64:", int64(math.MaxInt64))

		// Float32 is an exported function from rand
		fmt.Println("Random float32:", rand.Float32())
	}

	{
		// Invoke third party function Go() from quote package
		fmt.Println("Quote:", quote.Go())
	}

	{
		// Read weight and height from arguments if provided
		var weightKg float64 = 64
		var heightCm float64 = 167
		if len(os.Args) == 3 {
			var err error
			weightKg, err = strconv.ParseFloat(os.Args[1], 64)
			if err == nil {
				heightCm, err = strconv.ParseFloat(os.Args[2], 64)
			}
			if err != nil {
				fmt.Printf("Error: %s\nUsage: %s weight-in-kg height-in-cm\n", err, os.Args[0])
				return
			}
		}

		// Invoke the exported method BMI in bmi package
		b := bmi.BMI(weightKg, heightCm)
		fmt.Println("BMI =", b)

		// Get the guidelines for the calculated BMI from the nested package
		s := bmiguide.GuideBMI(b)
		fmt.Println("BMI guidelines:", s)
	}
}
