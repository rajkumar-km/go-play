/*
benchmark_echo measures the performance of Go code

Go standard lib "time" bundles with the time related utilities
  - time.Now() returns current time
  - time.Since(t) returns the duration from time t and now.

Also, we can write benchmark tests like we write test functions:
  - func BenchmarkSomeName(b *testing.B)
  - The b.N contains the argument to determine the number of runs
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// echo1 uses for/range loop to append command line arguments
func echo1() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// echo2 uses strings.Join to append command line arguments
func echo2() {
	fmt.Println(strings.Join(os.Args, " "))
}

// echo3 uses Println to append command line arguments
func echo3() {
	fmt.Println(os.Args)
}

// echo4 uses for/range loop and print the command line arguments with index
func echo4() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}

func main() {
	var durations []time.Duration

	// Measure echo1 running time
	start := time.Now()
	echo1()
	duration := time.Since(start)
	durations = append(durations, duration)

	// Measure echo2 running time
	start = time.Now()
	echo2()
	duration = time.Since(start)
	durations = append(durations, duration)

	// Measure echo3 running time
	start = time.Now()
	echo3()
	duration = time.Since(start)
	durations = append(durations, duration)

	// Measure echo4 running time
	start = time.Now()
	echo4()
	duration = time.Since(start)
	durations = append(durations, duration)

	// Print the results
	for i, dur := range durations {
		fmt.Printf("echo%d: %v\n", i+1, dur)
	}
}
