package example_test

import (
	"fmt"

	"github.com/rajkumar-km/go-play/go-basics/17-testing/example"
)

func ExampleBubbleSort() {
	v := []int{20, 40, 60, 100, 90, 50, 30, 70, 10, 80}
	example.BubbleSort(v)
	fmt.Printf("%#v", v)
	// Output:
	// []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
}