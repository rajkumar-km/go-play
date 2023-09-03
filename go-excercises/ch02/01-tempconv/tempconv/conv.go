package tempconv

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

// KToC converts temperature from Kelvin to Celsius
// C = K - 273.15
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK converts temperature from Celsius to Kelvin
// K = C + 273.15
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
