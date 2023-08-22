/*
echo_args_with_index prints all the command line arguments separated by space

os.Args is a built-in slice consists of all the command line arguments
os.Args[0] contains the executable name

Combination of "for" and "range" statements can be used to iterate charactors in the
strings, arrays, slices, and maps
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
