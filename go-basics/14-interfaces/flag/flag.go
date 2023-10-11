/*
flag demonstrates using the flag package for processing command line arguments
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// time.Duration is built into Go to pass duration values
	// So we can supply arguments like -d [13s | 2h30m | etc]
	duration := flag.Duration("d", 1*time.Second, "Duration")

	// Likewise we can implement flag.Value interface in our custom type and use it as args
	// Lets use the custom type Celsius and pass values like -t [37°C | 95°F]
	var temperature Celsius = 33 // 33 is the default value
	flag.CommandLine.Var(&temperature, "t", "Temperature")

	// flag.Parse parses os.Args based on the above flag registrations and populates
	// duration/temperature
	flag.Parse()
	fmt.Println(duration, temperature)
}

// Celsius holds the temperature in celsius unit
// We need to implement the flag.Value interface to support our types
// Let's implement flag.Value methods to pass this Celsius type as command line argument
// type Value interface {
// 		String() string
// 		Set(string) error
// }
type Celsius int

// String() returns the celsius value as string
func (t Celsius) String() string {
	return fmt.Sprintf("%d°C", t)
}

// Set() sets the temperature value in either celsius or fahrenheit (like 37°C or 95°F)
func (t *Celsius) Set(s string) error {
	var val int
	var unit string
	fmt.Sscanf(s, "%d%s", &val, &unit)
	// ignoring the Sscanf error since it will be taken care by the switch
	
	switch unit {
	case "C", "°C":
		// celsius
		*t = Celsius(val)
	case "F", "°F":
		// convert the input from fahrenheit to celsius
		*t = Celsius((val - 32) * 5 / 9)
	default:
		return fmt.Errorf("unsupported temperature unit: %s, use like 37°C or 95°F", s)
	}
	return nil
}
