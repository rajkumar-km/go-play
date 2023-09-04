package conv

import "fmt"

type Feet float64
type Meter float64

// String is the implementation of fmt.Stringer interface to use fmt.Print
func (v Feet) String() string {
	return fmt.Sprintf("%.2f feet", v)
}

func (v Meter) String() string {
	return fmt.Sprintf("%.2f meter", v)
}

// FeetToMeter converts the length from Feet to Meter
// M = F * 0.304
func FeetToMeter(f Feet) Meter {
	m := f * 0.304
	return Meter(m)
}

// MeterToFeet converts the length from Meter to Feet
// F = M / 0.304
func MeterToFeet(m Meter) Feet {
	f := m / 0.304
	return Feet(f)
}
