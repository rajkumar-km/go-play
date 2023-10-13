/*
sort demonstrates sorting range of values using sort.Interface

  - Generally, other languages supports sorting based on the data type and a compare functions.

  - Go's built in sort package is unusual. It does not know anything about data type and just
    provides a common interface to work with any sequence.

  - To sort any sequence, we generally need to implement these three functions and invoke
    sort.Sort()

    type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence
    Swap(i, j int)
    }

  - sort package uses hybrid sorting algorithms to provide better performance.

  - It has built in implementation for []int, []float64, and []string types.

  - sort.Ints(), sort.Strings(), sort.Float64s() are the built in functions.

  - Functions like sort.IntsAreSorted() can be used to check if the sequence is sorted.
*/
package main

import (
	"bytes"
	"fmt"
	"sort"
)

// Drive contains the storage drive details
type Drive struct {
	serial string
	capacity uint64
	freeSpace uint64
}

// DumpDrives returns the string dump of all drives
func DumpDrives(d []*Drive) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%-6s  %20s  %20s\n", "SERIAL", "CAPACITY", "FREE SPACE"))
	for i := range d {
		buf.WriteString(fmt.Sprintf("%-6s  %20d  %20d\n", d[i].serial, d[i].capacity, d[i].freeSpace))
	}
	return buf.String()
}

// DriveListByFreeSpace contains methods to sort drives by free space
type DriveListByFreeSpace []*Drive
func (d DriveListByFreeSpace) Len() int { return len(d) }
func (d DriveListByFreeSpace) Less(i, j int) bool {	return d[i].freeSpace > d[j].freeSpace }
func (d DriveListByFreeSpace) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

// DriveList is an optimized version to support sort by any column by providing the less function
type DriveList struct {
	drives []*Drive
	less func(x, y *Drive) bool
}
func (d DriveList) Len() int { return len(d.drives) }
func (d DriveList) Less(i, j int) bool { return d.less(d.drives[i], d.drives[j]) }
func (d DriveList) Swap(i, j int) { d.drives[i], d.drives[j] = d.drives[j], d.drives[i] }

func main() {
	// List of drives
	drives := []*Drive{
		{serial: "X0", capacity: 42949672960, freeSpace: 2333556},
		{serial: "B0", capacity: 107374182400, freeSpace: 1024},
		{serial: "A0", capacity: 10737418240, freeSpace: 23423566},
	}

	// convert []*Drive to DriveListByFreeSpace named type since it has the sort.Interface
	// implemented to sort the drives by free space.
	sort.Sort(DriveListByFreeSpace(drives))
	fmt.Println("Sort by free space")
	fmt.Println(DumpDrives(drives))

	// We don't have to define a new type for reverse sort.
	// sort.Reverse converts it to sort.reverse type which intern compose the original data
	// (sort.Interface). It has the the overriden Less() method which invokes the original Less()
	// method with indexes swapped Less(j, i)
	sort.Sort(sort.Reverse(DriveListByFreeSpace(drives)))
	fmt.Println("Reverse sort")
	fmt.Println(DumpDrives(drives))

	// So, we need to write a new named type to sort by a separate column.
	// This can be optimized further by having a dynamic function value for "less"
	dl := DriveList{
		drives: drives,
		less: func(x, y *Drive) bool {
			return x.capacity > y.capacity
		},
	}
	sort.Sort(dl)
	fmt.Println("Sort by capacity")
	fmt.Println(DumpDrives(drives))
	
	// The above examples are given to demonstrate sort.Interface
	// But, Go already has the built-in methods to sort slices with a function
	sort.Slice(drives, func(i, j int) bool { return drives[i].serial < drives[j].serial })
	fmt.Println("Sort by serial")
	fmt.Println(DumpDrives(drives))
	fmt.Println("sort.IsSorted = ", sort.IsSorted(dl))
}
