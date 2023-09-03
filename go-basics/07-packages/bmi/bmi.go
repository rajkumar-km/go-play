/*
Package bmi calculates the Body Mass Index (BMI) from height and weight
*/
package bmi

// BMI calculates the body mass index
// Since it starts with an upper case letter, it is accessible by other packages
func BMI(weightKg float64, heightCm float64) float64 {
	heightM := cmToM(heightCm)
	// Formula (weight in kg / height in metre squared)
	return weightKg / (heightM * heightM)
}

// cmToM converts the value in centimetres to metre
// Since it starts with lower case letter, it is private to current package
func cmToM(cm float64) float64 {
	return cm / 100
}
