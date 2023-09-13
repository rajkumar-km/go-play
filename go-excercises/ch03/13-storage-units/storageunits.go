/*
storageunits demonostrates the defining constants for storage units such as KB and MB
*/
package main

import "fmt"

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB // overflows 64-byte
	YB = 1000 * ZB // overflows 64-byte
)

func main() {
	for _, v := range []uint64{KB, MB, GB, TB, PB, EB} {
		fmt.Printf("%d\n", v)
	}
}
