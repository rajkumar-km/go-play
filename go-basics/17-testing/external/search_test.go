/*
search_test demonstrates writing external test packages

  - Generally, the test package is same as the original package for whitebox testing. So that it
    can access the internal members.
  - But sometimes this can create import cycles and caused build failures. For example, net/http
    uses the net/url package. Test functions of net/url package uses net/http and form the import
    cycle.
  - So, we can fix this by writing the test in external package net/url_test. The test can import
    both net/url and net/http packages.
  - However, we lost the ability to access the internal functions while writing tests in external
    packages. There is a trick to allow internal access for white box testing.
  - Write a test file (typically named "export_test.go") in the same package as the original and
    expose the internal members outside since it has the access. Anyway the "_test.go" is not part
    of the production build and applicable only to test.
  - go list tool is handy to differenciate these files:
    go list -f={{.GoFiles}} fmt       # list production files included in go build
    go list -f={{.TestGoFiles}} fmt   # list test files in same package
    go list -f={{.XTestGoFiles}} fmt  # list test files in external test package
*/
package search_test

import (
	"testing"

	search "github.com/rajkumar-km/go-play/go-basics/17-testing/external"
)

func TestBinarySearch(t *testing.T) {
	v := []int{20, 30, 40, 50, 70, 100}
	if ! search.IsSorted(v) {
		// Exit the test if when using invalid test data
		t.Fatalf("Input array not sorted for binary search: %v", v)
	}

	idx := search.BinarySearch(v, 0, len(v)-1, 70)
	want := 4
	if idx != want {
		t.Errorf("search.BinarySearch(v, %d) = %d, want %d", 70, idx, want)
	}
}

func TestIsSorted(t *testing.T) {
	v := []int{20, 30, 40, 50, 70, 100}
	if ! search.IsSorted(v) {
		t.Errorf("search.IsSorted(v) = false")
	}

	v2 := []int{30, 40, 50, 10, 20}
	if search.IsSorted(v2) {
		t.Errorf("search.IsSorted(v2) = true")
	}
}