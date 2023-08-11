/*
mypkg is a sample user defined package
*/
package mypkg

import (
	"fmt"
	// using alias name for packages to give a shorter name or
	// Avoid name conflicts if multiple packages presents with the same name
	str "strings"
)

// MyFunc starts with upper case letter. So, it is accessible by other packages
func MyFunc() {
	fmt.Println("My Func")
}

// myPrivateFunc starts with lower case letter, so it is private to current package
func myPrivateFunc() {
	fmt.Println("My Private Func")
	fmt.Println(str.ToUpper("hello"))
}
