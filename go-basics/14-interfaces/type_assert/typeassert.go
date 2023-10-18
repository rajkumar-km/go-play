/*
typeassert demonstrates Go type assertions
  - Assertain the actual type stored in an interface
  - Assertion also performs convertion to the target type
*/
package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

func main() {
	demoTypeAssert()
	demoErrorDifferenciation()
	demoTypeAssertForInterface()
}

// demoTypeAssert shows how to assertain the actual type stored in an interface.
// An interface{} variable is capable of storing data of any type. Go provides
// few ways to determine the actual type stored in the variable.
// Example to use when you are sure about the underlying type:
//
//	var data interface{} = 10
//	myInt := data.(int) // Note: If the data is not int, then it causes panic.
//
// Example to assert the type without causing panic
//
//	myInt, ok := data.(int)
//
// - Not only it asserts the concrete type, but also returns the concrete type on success.
func demoTypeAssert() {
	var data any = 10 // any is an alias of interface{}

	// Get the data as integer when you are sure about the underlying type.
	myInt := data.(int)
	fmt.Printf("data.(int) = %d\n", myInt)

	// panic: interface conversion: interface {} is int, not float32
	// myFloat := data.(float32)

	// Safe assertion without causing panic
	myFloat, ok := data.(float32)
	if ok {
		fmt.Printf("data.(float32) = %f\n", myFloat)
	} else {
		fmt.Printf("data is not float32\n")
	}
}

// demoErrorDifferenciation shows how to handle path errors through error
// interface
func demoErrorDifferenciation() {
	_, err := os.Open("/invalid/path")
	fmt.Println(err.Error()) // open /invalid/path: no such file or directory
	if os.IsNotExist(err) {
		fmt.Println("NotExist")
	}

	// How can we replicate os.IsNotExist()?
	// Does os.IsNotExist() checks if the returned error is: no such file or directory? No
	if strings.Contains(err.Error(), "no such file or directory") {
		// This seems ugly and it is not platform independent
	}

	// Actually os.Open() returns *fs.PathError wrapped in error interface
	// &fs.PathError{Op:"open", Path:"/invalid/path", Err:0x2}
	// So, this contains the operation, file path, and the error code.
	if pathErr, ok := err.(*fs.PathError); ok {
		fmt.Printf("%#v\n", pathErr)
		if pathErr == os.ErrNotExist {
			fmt.Println("os.ErrNotExist")
		}
	}
}

// demoTypeAssertForInterface demonstrates converting one interface to another
func demoTypeAssertForInterface() {

	// Type assertion also works to check against an interface. If the underlying type satisfies
	// the interface then it returns the asserted interface. The original concrete type is
	// preserved.
	var temp Temperate = Celsius(10)

	// temp is the variable of Temperate interface
	// It can be converted to Stringer interface
	t, ok := temp.(fmt.Stringer)
	if ok {
		fmt.Println(t.String())
	} else {
		fmt.Printf("%T is not compatible with fmt.Stringer\n", temp)
	}

	// Another example: io.Writer interface ensures that it has Write() method which accepts
	// []byte slice. 
	// - But, if we need to write a string then a conversion is required like
	//   []byte(str) which is inefficient since it creates a copy.
	// - Instead, the WriteString() method available in certain concrete types which is efficient
	//   to write strings.
	// - So, we could improve write like:
	tempws, ok := temp.(interface{WriteString(io.Writer, string) (int, error)})
	if ok {
		tempws.WriteString(os.Stdout, "mystr")
	}
}

// Celsius records the temperate in celsius unit
type Celsius float64
type Temperate interface {
	// Celsius returns the temperate in celsius unit
	Celsius() Celsius
}

func (c Celsius) Celsius() Celsius {
	return c
}
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}
