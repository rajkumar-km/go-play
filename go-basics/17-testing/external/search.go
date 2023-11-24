/*
search provide functionalities to search elements in an array

Go allows to have different package name and need not be same as the folder name.
But import must point to the folder name.
*/
package search

// BinarySearch
func BinarySearch(v []int, low int, high int, x int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if v[mid] == x {
		return mid
	} else if v[mid] > x {
		return BinarySearch(v, 0, mid-1, x)
	}
	return BinarySearch(v, mid+1, high, x)
}

func isSorted(v []int) bool {
	for i := 1; i < len(v); i++ {
		if v[i] < v[i-1] {
			return false
		}
	}
	return true
}
