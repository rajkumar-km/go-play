package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

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

// CToF converts temperature from Celsius to Fahrenheit
// F = (9/5)C + 32
func CToF(c Celsius) Fahrenheit {
	f := (9.0/5)*c + 32
	return Fahrenheit(f)
}

// FToC converts temperature from Fahrenheit to Celsius
// C = F - 32 * (5/9)
func FToC(f Fahrenheit) Celsius {
	c := (f - 32) * (5.0 / 9)
	return Celsius(c)
}
