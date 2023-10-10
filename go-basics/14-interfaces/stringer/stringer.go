/*
stringer demonstrates the built-in fmt.Stringer interface
*/
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type IntSet struct {
	values []int
}

// String returns the string representation of IntSet
func (set *IntSet) String() string {
	var b bytes.Buffer
	b.WriteByte('{')
	for i,v := range set.values {
		b.WriteString(strconv.Itoa(v))
		if (i < len(set.values)-1) {
			b.WriteByte(',')
		}
	}
	b.WriteByte('}')
	return b.String()
}

func main() {
	var set IntSet = IntSet{values: []int{10, 20, 30}}
	// The following statement will cause compile error, because the String() method is defined as
	// pointer receiver type. But the assignment here is value type. Any value assigned to an
	// interface variable is not addressable, so compiler can not invoke pointer receiver method.
	// So, prefer to define the String() method as valye type that works for both.
	// var _ fmt.Stringer = set  
	var _ fmt.Stringer = &set // compatible with fmt.Stringer and works

	// fmt methods accepts any type of arguments using interface{} type and checks if it satisfies
	// fmt.Stringer method. If so, invokes the String() method, otherwise prints raw data.
	fmt.Println(set) // does not invoke String()
	fmt.Println(&set) // invoke String()
}