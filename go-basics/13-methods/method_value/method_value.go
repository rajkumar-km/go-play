/*
method_value demonstrates using method as values in Go

Like function values, methods can also be considered as values in Go
- obj.Method() is a method call
- obj.Method is a method value. This binds both object and method together.

  - A method value can be passed as an argument for function callbacks
  - It can also be stored in a variable and called later.

m := obj.Method
m()
*/
package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func (c *Counter) Incr() {
	c.count++
}

func (c *Counter) Print() {
	fmt.Println(c.count)
}

func main() {
	c := Counter{}

	// Store method value in a variable, and call it later
	// Method value binds receiver argument and method together
	pr := c.Print
	pr()

	// Also, we can pass method value to call backs which is useful
	time.AfterFunc(100*time.Millisecond, c.Incr)
	time.Sleep(101 * time.Millisecond)

	pr()
}
