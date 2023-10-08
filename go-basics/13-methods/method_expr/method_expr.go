/*
method_expr demonstrates using methods as expression in Go

- Go allows to refer the methods as "TypeName.Method" or "(*TypeName).Method"
- Method expressions can be passed as argument to call backs or
- Assigned to variables and called later but with receiver arguments.
*/
package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Incr() {
	c.count++
}

func (c *Counter) Decr() {
	c.count--
}

func (c Counter) Print() {
	fmt.Println(c.count)
}

func main() {
	incr := false

	var op func(*Counter)
	if incr {
		op = (*Counter).Incr // method expression
	} else {
		op = (*Counter).Decr // method expression
	}

	counters := []*Counter{{count: 1}, {count: 2}}
	for _, c := range counters {
		op(c) // call method expression with pointer receiver
	}

	pr := Counter.Print // value receiver method expression
	for _, c := range counters {
		pr(*c) // provide value type for value receiver method
	}
}
