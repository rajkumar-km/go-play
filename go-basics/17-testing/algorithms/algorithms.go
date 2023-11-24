/*
Package algorithms provides the utility functions for key algorithms
*/
package algorithms

// LinearSearch searches for the element e in the slice v
// Returns the matching index, -1 otherwise.
func LinearSearch(v []int, e int) int {
	for i := range v {
		if v[i] == e {
			return i
		}
	}
	return -1
}
