package main

import (
	"fmt"
	"time"
)

// DemoVariadic demonstrates variable length arguments to Go functions
//   - A variable length argument can be specified with "..." before the type
//   - Example: "numbers ...int"
//   - Inside the function, the type becomes a slice ([]numbers)
//
// In addition, we can also pass a slice with "..." to expand it as a
// argument list while calling the function. Example: sumAll(numbers...)
//   - The fmt.Printf and other formatting functions are the great examples for
//     variadic functions.
func DemoVariadic() {
	sum := sumAll(1, 3, 4)
	fmt.Printf("sumAll(1, 3, 4) = %d\n", sum)

	nums := []int{10, 20, 30}
	sum = sumAll(nums...)
	logf("sumAll(nums...) = %d\n", sum)
}

// sumAll sums all the numbers passed to it and returns the sum
func sumAll(numbers ...int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// logf formats and logs the message with timestamp
// format indicates the fmt format string
// args accepts variable length of arguments. interface{} indicates any type
func logf(format string, args ...interface{}) {
	format = time.Now().Format("2006-01-02 15:04:05 ") + format
	fmt.Printf(format, args...)
}