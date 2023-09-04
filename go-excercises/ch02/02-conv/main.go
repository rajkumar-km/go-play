/*
conv is a general purpose unit conversion utility that supports:
- length: feet and meters
- mass: kilogram and pounds
- temperature: celsius and fahrenheit

Sample input:

	101.2f
		- fahrenheit to celsius
	100c
		- celsius to fahrenheit
	64.2kg
		- kilogram to pound
	130.5lb
		- pound to kilogram
	67'
		- feet to meter
	1.5m
		- metre to inch
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rajkumar-km/go-play/go-excercises/ch02/02-conv/conv"
)

func main() {
	// Read input from command line arguments if provided
	if len(os.Args) > 1 {
		for _, s := range os.Args[1:] {
			fmt.Printf("%s = %s\n", s, convert(s))
		}
		return
	}

	// Otherwise read from stdin and process line by line
	fmt.Println(`Enter the values in the format: <value><unit>. Eg: 37c, 40', or 10kg`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s = %s\n", scanner.Text(), convert(scanner.Text()))
	}
}

// convert interprets the input string in unit and invokes the matching
// conversion function
func convert(s string) string {
	val, unit, err := parseInput(s)
	if err != nil {
		return err.Error()
	}

	switch unit {
	case "c": // Celsius
		return conv.CToF(conv.Celsius(val)).String()
	case "f": // Fahrenheit
		return conv.FToC(conv.Fahrenheit(val)).String()
	case "'": // Feet
		return conv.FeetToMeter(conv.Feet(val)).String()
	case "m": // Meter
		return conv.MeterToFeet(conv.Meter(val)).String()
	case "kg": // Kilogram
		return conv.KGToLB(conv.Kilogram(val)).String()
	case "lb": // Pound
		return conv.LBToKG(conv.Pound(val)).String()
	default:
		return fmt.Sprintf(`%s unit is unsupported. Try one of these: 10.2f, 37c, 50', 3.5m, 10kg, 30lb`, unit)
	}
}

// inputExp represents the user input string
// Example: 10.2f, -37.1c, 50', 3.5m, 10kg, 30lb
var inputExp *regexp.Regexp = regexp.MustCompile(`(-?[0-9.]+)([a-z']+)`)

// parseInput parses the user input based on inputExp
func parseInput(s string) (float64, string, error) {
	m := inputExp.FindStringSubmatch(s)
	if len(m) != 3 {
		err := fmt.Errorf(`Unsupported format. Use format <float-value><unit>. Eg: 37c`)
		return 0, "", err
	}

	val := m[1]
	unit := m[2]

	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, "", err
	}

	return v, unit, nil
}
