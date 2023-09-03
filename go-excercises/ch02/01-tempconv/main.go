package main

import (
	"fmt"

	"github.com/rajkumar-km/go-play/go-excercises/ch02/01-tempconv/tempconv"
)

// Demonstrates the temperature conversions with a Go package
func main() {
	fmt.Printf("Absolute Zero or Kelvins Zero is %.2f°C\n", tempconv.AbsoluteZeroC)
	fmt.Printf("%s = %s\n", tempconv.FreezingC, tempconv.CToF(tempconv.FreezingC+1))
	fmt.Printf("%s = %s\n", tempconv.Fahrenheit(100), tempconv.FToC(100))
	fmt.Printf("%s = %s\n", tempconv.AbsoluteZeroC, tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Printf("%s = %s\n", tempconv.Kelvin(0), tempconv.KToC(0))
}
