package profile_test

import (
	"testing"

	"github.com/rajkumar-km/go-play/go-basics/17-testing/benchmark"
)

func BenchmarkBinarySearch(b *testing.B) {
	v := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < b.N; i++ {
		benchmark.BinarySearch(v, 0, len(v)-1, 80)
	}
}