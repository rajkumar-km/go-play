/*
Package algorithms_test demonstrates testing packages in Go

  - Go has the inbuilt support for testing the code. This is no different from writing the
    functional code.
  - Test code resides the same directory but the filenames ending with "_test.go". This can
    contain testing, benchmark, and example code. The "go build" skip files ending with "_test.go"
  - The "go test" command can be used to run the test code. This would build both the actual
    code and the test code. Create an test executable automaticaly to invoke test functions.
  - Go tool allows to have two packages in the same directory to accomodate test code. So, we have
    "algorithm" and "algorithm_test" packages in the same path.

Commands to run Test:
  - go test
  - go test -v # verbose mode to print test function names and its results
  - go test -run="NotFound|Invalid" # run only matching test functions

Test Functions:
  - Each test file imports the "testing" package from standard library and have functions starting
    with keyword "Test" and accepts the testing.T argument. The arguments is useful to report the
    errors.
  - Use t.Error or t.Errorf to display error and continue to run further tests
  - Use t.Fatal or t.Fatalf to break execution
  - It is usually a good practice to write the test first and validate the errors reported. Later,
    the functionality or fix can be implemented to ensure that we address the right problem.
*/
package algorithms_test

import (
	"testing"

	"github.com/rajkumar-km/go-play/go-basics/17-testing/algorithms"
)

func TestLinearSearch(t *testing.T) {
	v := []int{10, 40, 30, 20, 50}
	idx := algorithms.LinearSearch(v, 20)
	if idx != 3 {
		t.Error("LinearSearch(v, 20) != 3")
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	v := []int{10, 40, 30, 20, 50}
	idx := algorithms.LinearSearch(v, 100)
	if idx != -1 {
		t.Error("LinearSearch(v, 100) != -1")
	}
}

func TestLinearSearchComprehensive(t *testing.T) {
	v := []int{10, 40, 30, 20, 50}

	// Use this comprehensive table-driven model to cover bunch of various inputs
	var cases = []struct{
		input int
		want int
	}{
		{10, 0},
		{40, 1},
		{30, 2},
		{20, 3},
		{50, 4},
		{60, -1},
		{100, -1},
		{-1, -1},
	}

	for _,c := range cases {
		idx := algorithms.LinearSearch(v, c.input)
		if idx != c.want {
			// Usual form of error message includes "want", but skip this if want is a boolean
			t.Errorf("LinearSearch(v, %d) == %d, want %d", c.input, idx, c.want)
		}
	}
}