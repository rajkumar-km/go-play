package main

import "fmt"

// Package variables initialized as per the order except that
// the dependencies are resolved first.
var a = 10     // initialized first
var b = c - 10 // initialized third because it has dependency of c
var c = f()    // initialized second after calling function f
func f() int   { return a + 10 }

// DisplayName is a package level variable
var DisplayName string = initDisplayName()

// initDisplayName initializes the global variable DisplayName
// Global variables are initialized even before the init function
// So, the initDisplayName() is called even before init()
func initDisplayName() string {
	fmt.Println("package variable initialization first")
	return "Variables"
}

// init function is executed by default even before the main function.
func init() {
	fmt.Println("init() function executed after package variables initialization")
}

// There can be more than one init functions for the same package
// and they are executed in the order defined.
func init() {
	fmt.Println("multiple init() functions are executed in order")
}
