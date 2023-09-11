package main

import (
	"fmt"
)

// iota is a Go's predeclared const that helps to auto initialize the constants
//   - iota is set to zero whenever a const block is started
//   - iota is auto incremented by 1 on every new line
//   - The first constant in the const block must be initialized.
//   - Subsequent constants can be left uninitialized. Go by default applies the
//     last expression to subsequent constants if left empty.
//
// iota is powerful with expresssion, use it only for necessary cases
//   - Use the hard coded values instead if the constant values are exported or
//     serialized by consumers. Otherwise, it can cause trouble if a new constant
//     is introduced in between the order.
//   - Ensure that the code is readable and understandable
const (
	Fatal = iota // iota = 0
	Error        // iota = 1
	Warn         // iota = 2
	Info         // iota = 3
	Debug        // iota = 4
	_            // iota = 5, use _ to skip a value
	Trace        // iota = 6
)

// iota starts from zero and always incremented by one for each newline
// But it can be combined with expressions to achieve various results
// For example, using the expression "iota + 1" in the first constant
// is applied for the subsequent constants as well.

// Day is like a c++-style enum
type Day int

const (
	Sunday    Day = iota + 1 // iota=0 + 1 = 1
	Monday                   // iota=1 + 1 = 2
	Tuesday                  // iota=2 + 1 = 3
	Wednesday                // iota=3 + 1 = 4
	Thursday                 // iota=4 + 1 = 5
	Friday                   // iota=5 + 1 = 6
	Saturday                 // iota=6 + 1 = 7
)

// Another good example is doing bitwise operations with iota
const (
	_   = 1 << (iota * 10) // 1 << (iota=0 * 10) = 0
	KiB                    // 1 << (iota=1 * 10) = 1024
	MiB                    // 1 << (iota=2 * 10) = 1048576
	GiB                    // 1 << (iota=3 * 10) = 1073741824
	TiB                    // 1 << (iota=4 * 10) = 1099511627776
)

// If iota is used outside the const block, then it will always be zero
const zeroInt int = iota       // 0
const zeroFloat float32 = iota // 0

// DemoIotaSamples demonstrates the various usages of iota
func DemoIota() {
	fmt.Println(Fatal, Error, Warn, Info, Debug, Trace)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(KiB, MiB, GiB, TiB)
}
