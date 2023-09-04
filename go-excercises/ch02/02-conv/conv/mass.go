package conv

import "fmt"

type Kilogram float64
type Pound float64

// String is the implementation of fmt.Stringer interface to use fmt.Print
func (v Kilogram) String() string {
	return fmt.Sprintf("%.2f kg", v)
}

func (v Pound) String() string {
	return fmt.Sprintf("%.2f lb", v)
}

// KGToLB converts the mass from Kilogram to Pound
// LB = KG * 2.204
func KGToLB(kg Kilogram) Pound {
	lb := kg * 2.204
	return Pound(lb)
}

// LBToKG converts the mass from Pound to Kilogram
// KG = LB / 2.204
func LBToKG(lb Pound) Kilogram {
	kg := lb / 2.204
	return Kilogram(kg)
}
