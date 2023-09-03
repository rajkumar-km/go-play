/*
Package tempconv converts the temperate between different scales:
  - Celsius
  - Fahrenheit
  - Kelvin
*/
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// Kelvin scale defines this as absolute zero temperature
	// However it is difficult to reach zero kelvins as it requires infinite energy
	// Also, atoms and molecules would still have some irreducible motion
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// String is the implementation of fmt.Stringer interface to use fmt.Print
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f°K", k)
}
