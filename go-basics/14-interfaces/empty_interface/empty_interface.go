/*
empty_interface demonstrates the use of interface{} in Go

  - The empty interface{} has no methods, and all the types satisfy this interface. So,
    any value can be assigned to a variable of type interface{}.
  - Empty interface{} has an alias named "any"
  - But, we can not perform any operations on it since it does not have methods. The why
    do we have it? It can be passed as common interface{} and converted to its original type
    through type assertions or type switch wherever required.
  - For example, a generic Stack can be implemented to push/pop any type of values and it
    can be converted and processed by the consumer after pop()
  - The fmt.Print family of functions accepts any values as arguments. This is because it
    has interface{} type as arguments.
*/
package main

import (
	"fmt"
)

func main() {
	var data any // or "var data interface{}"

	// Any type of values can be assigned to it
	data = false
	data = 10
	data = 12.4
	data = 1 + 3i	
	data = 'X'
	data = "hello"
	data = []int{1,2,3}
	data = map[string]int{}
	data = struct{}{}	

	Println(data)
}

// fmt.Println uses the following format to accept any type and any number of arguments
func Println(v ...any) (int, error) {
	return fmt.Println(v...) // v... expands the slice to list of arguments
}