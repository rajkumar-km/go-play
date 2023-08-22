/*
echo_args prints all the command line arguments separated by space

os.Args is a built-in slice consists of all the command line arguments
os.Args[0] contains the executable name

strings.Join(slice, sep) joins the given array/slice with separator
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {	
	fmt.Println(strings.Join(os.Args, " "))	
}
