/*
typedeclaration demonstrates defining new named types from an existing Go type

The built-in types such as int are used for various purposes. For example, an int type can be
used to store a loop index, or an epoch timestamp, or timeout seconds. The named type allow
us create a new type and separate its uses so that it can't be mixed unintentionally.

	type name underlying-type
	type Celsius float64
	type Farenheit float64

We can declare constants/variables with the new named type. But, we can not compare or perform
operations between a Celsius and Farenheit variable although the underlying type is float64.

	var c Celsius
	var f Farenheit
	c == f // compile error
	c + f  // compile error

However, an explicit type conversion is allowed when the underlying type is same

	c + Celsius(f) // allowed to mix types with explicit typecast

In addition, the advantages of named types are more with more complex types such as struct.
Names types can be added with new methods like below

	func (c Celsius) String() string {
		return fmt.Sprintf("%g°C", c)
	}
	var c Celsius = 100
	fmt.Println(c)          // "100°C"
	fmt.Println(c.String()) // "100°C"

The String() method is often added the user defined types.
Have you thought of how fmt.Print works when passing any type of arguments?
  - It uses interface{} type of argument to accept any values
  - Checks whether the given argument is compatible with fmt.Stringer interface
    (must have the method String()). If so, executes the String() method
  - Otherwise, uses reflection features. See 17-reflection for more details.
*/
package main

import "fmt"

// Celsius stores the temperature in celsius unit
type Celsius float64

// Farenheit stores the temperature in farenheit unit
type Farenheit float64

// constants/variables can be declared with new named type
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius     = 0
	BoilingC Celsius      = 100
)

// String() method is added to be compatible with fmt.Stringer interface so that
// the fmt print functions use this method to print celsius values
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// DemoTypeDeclaration demonstrates defining new types from an existing Go type
func DemoTypeDeclaration() {
	var c Celsius = FreezingC
	var f Farenheit = 102

	fmt.Println(c >= 0)           // allowed
	fmt.Println(c <= BoilingC)    // allowed
	fmt.Println(f == 0)           // allowed
	// fmt.Println(c == f)        // compile error
	// fmt.Println(f == BoilingC) // compile error
	// fmt.Println(f - c)         // compile error
	fmt.Println(f - Farenheit(c)) // allowed with explicit type cast

	c = BoilingC
	fmt.Println(c)          // 100°C
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%s\n", c)   // 100°C
	fmt.Printf("%v\n", c)   // 100°C
	fmt.Println(float64(c)) // 100; does not call String()
	fmt.Printf("%g\n", c)   // 100; does not call String()
}