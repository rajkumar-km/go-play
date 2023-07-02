/*
mynestedpkg is a sample nested package under mypkg/
*/
package mynestedpkg

import "fmt"

// MyNestedPkgFunc is accessible from other packages since it starts with upper case
func MyNestedPkgFunc() {
	fmt.Println("My Nested Pkg Func")
}
