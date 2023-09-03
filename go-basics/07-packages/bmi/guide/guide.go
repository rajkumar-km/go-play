/*
Package guide provide guidelines for body mass index (BMI)
It is nested under the bmi package
*/
package guide

import "fmt"

const (
	MinHealthyBMI = 18.5
	MaxHealthyBMI = 24.9
	MinObeseBMI   = 30.0
)

// GuideBMI provide the guidelines for the given BMI
func GuideBMI(b float64) string {
	var g string

	switch {
	case b < MinHealthyBMI:
		g = `You are underweight as per your BMI. Eat more calories than you burn.`

	case b <= MaxHealthyBMI:
		g = "Your weight is ideal. Stay healthy"

	case b < MinObeseBMI:
		g = "You are overweight as per your BMI. Burn more calories than you eat."

	default:
		g = "Your BMI indicates obesity. Consult a doctor"
	}

	fmt.Printf("Healthy BMI ranges between %g and %g\n", MinHealthyBMI, MaxHealthyBMI)
	return g
}
