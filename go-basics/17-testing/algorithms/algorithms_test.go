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
    errors and log debug messages.
  - Use t.Error or t.Errorf to display error and continue to run further tests
  - Use t.Fatal or t.Fatalf to break execution
  - It is usually a good practice to write the test first and validate the errors reported. Later,
    the functionality or fix can be implemented to ensure that we address the right problem.

Table-driven tests:
  - A comprehensive table-driven model to cover bunch of various inputs
  - See TestLinearSearchTableDriven() for example

Randomized tests:
  - Inputs can be generated in random for testing instead of using a table.
  - But, how do we know the expected output. There are two ways:
    1. Write an alternative implementation which is straightforward and may be less efficient to
    produce the expected output.
    2. Generate the inputs in a such a way that we can guess the output
  - It may be difficult to debug the failing cases when using random inputs. Instead of dumping the
    whole lot of information, simply log the random seed or the input which is sufficient to
    reproduce the failure again.
  - See TestLinearSearchRandom() for example

Writing Effective Tests:
  - Many newcomers go Go surprised by the minimalism of Go's testing framework. It does not have
    the following features: Setup and Teardown hooks for performing certain operations before and
    after the test runs. Utility libraries for assertions, formatting, aborting a failed test.
  - Although these helpers make the tests very concise, the resulting code often seems like foreign
    language and not like Go. Also, the error messages provided by the assert functions may not
    have the sufficient context.
  - Go expects test authors to do most of the works, defining functions to avoid repetition.
  - A good test should not break and report the problem with context. Developers should not look
    at the code to debug a test failure.
  - Also, it should return immediately after the first failure and but provide list of failures.

Avoiding Brittle Tests:
  - An application that fails when entering valid new input is buggy.
  - A test that fails when a sound change was made to the program is called brittle.
  - To write easily maintainable tests:
  - Check only the properties that you care about.
  - Test stable API interfaces more often than the interface functions.
  - Do not match the full error message, instead match only key sub string that won't change.
  - If needed write functions to parse the complex output to its essence so that the assertions
    will be reliable.

Coverage:
  - By its nature, testing is never complete.
  - "Testing shows the presence of bugs and not the absense" - Edsger Dijkstra
  - The degree in which test suite exercise the package is called the test coverage.
  - Go test has the built-in options to measure the code coverage:
  - go tool cover  # shows help
  - go test --cover # returns the coverage summary in percent
  - go test --coverprofile=c.out # generates the log file with coverage report
    Go test modifies the copy of a source code for testing. It set a boolean variable for every
    block to track the coverage. Finally writes the coverage flag in the log file c.out
  - go test --coverprofile=c.out --covermode=count
    Instead of having boolean variable for every block, a counter is used to track how many times
    a block is covered in the test.
  - go tool cover --html=c.out
    Generates a html output from the cover profile generated earlier and open in it browser.
  - Achieving 100% coverage may seems like a goal, but it is not the usual case:
  - A line is bug free not just because a statement is executed. Some statements are never
    reachable: Say a panic() writen in a unreable block.
  - Instead we should focus on repeated test in a complex path with different inputs.
*/
package algorithms_test

import (
	"math/rand"
	"testing"
	"time"

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

func TestLinearSearchTableDriven(t *testing.T) {
	v := []int{10, 40, 30, 20, 50}

	// Use this comprehensive table-driven model to cover bunch of various inputs
	var cases = []struct {
		input int
		want  int
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

	for _, c := range cases {
		idx := algorithms.LinearSearch(v, c.input)
		if idx != c.want {
			// Usual form of error message includes "want", but skip this if want is a boolean
			t.Errorf("LinearSearch(v, %d) == %d, want %d", c.input, idx, c.want)
		}
	}
}

func TestLinearSearchRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		seed := time.Now().UTC().UnixNano()
		t.Logf("Random seed: %d", seed)
		rng := rand.New(rand.NewSource(seed))

		v := make([]int, 100)
		want := rng.Intn(100)
		v[want] = 30
		idx := algorithms.LinearSearch(v, 30)
		if idx != want {
			t.Errorf("LinearSearch(v, %d) == %d, want %d", 30, idx, want)
		}
	}
}
