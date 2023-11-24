/*
Package benchmark defines a sample function to experiment Go benchmarking
*/
package benchmark

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

// BinarySearch perform the binary search using recursive method
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

// BinarySearchNonRecursive is a non recursive version of binary search
func BinarySearchNonRecursive(v []int, low int, high int, x int) int {
	for low < high {
		mid := (low + high) / 2
		if v[mid] == x {
			return mid
		} else if v[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
