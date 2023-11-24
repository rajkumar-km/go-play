/*
benchmark_test demonstrates benchmarking Go programs

  - Benchmarking is a practice of measuring the performance on a fixed workload

  - Go benchmark functions start with Benchmark keyword and accepts *testing.B argument. This
    would provide the same methods as testing.T plus some extra for performance measurements.

  - testing.B also has a fixed number N to run the benchmark N times which can be controlled
    while running.

  - go test --bench=.      # dot matches all benchmarks in the package

  - go test --bench=Binary # runs only matching benchmark with the keyword
    goos: windows
    goarch: amd64
    pkg: github.com/rajkumar-km/go-play/go-basics/17-testing/benchmark
    cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
    BenchmarkBinarySearch-8                 295212319                3.492 ns/op
    BenchmarkBinarySearchNonRecursive-8     778451426                1.510 ns/op

  - The suffix 8 indicates the value of GOMAXPROCS which is to measure concurrent benchmarks.

  - The value 295212319 and 778451426 indicates the number of times the test is run. This is
    automatically determined by the Go. First it runs N for smaller count and understand the time
    taken. Based on that, it can prepare a count that is sufficient for stable measurements.

  - Every benchmark function implements a loop for N times. The loop is left to the benchmark
    functions instead of test driver, because we can have some initialization statements outside
    the loop.

  - The testing.B also provides methods to stop, resume, and reset timers in between, but this is
    rarely required.

  - Generally the fastest version allocates fewer memory. To measure the memory:

  - go test --bench=Binary --benchmem # measure memory usage

  - The benchmarks provided here are absolute. Relative benchmarks are interesting such as what
    if the function is called 100 times vs 1000 times. What is best buffer size? Which algorithm
    performs better? Just write different Benchmark functions which may call a single helper
    function with different arguments.
*/
package benchmark_test

import (
	"testing"

	"github.com/rajkumar-km/go-play/go-basics/17-testing/benchmark"
)

func BenchmarkLinearSearch(b *testing.B) {
	v := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < b.N; i++ {
		benchmark.LinearSearch(v, 80)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	v := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < b.N; i++ {
		benchmark.BinarySearch(v, 0, len(v)-1, 80)
	}
}

func BenchmarkBinarySearchNonRecursive(b *testing.B) {
	v := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < b.N; i++ {
		benchmark.BinarySearchNonRecursive(v, 0, len(v)-1, 80)
	}
}

// benchmarkBSearch performs relative benchmark testing with various size
func benchmarkBSearch(b *testing.B, size int) {
	v := make([]int, size)
	for i := 0; i < len(v); i++ {
		v[i] = i * 10
	}
	e := v[size%20]

	for i := 0; i < b.N; i++ {
		benchmark.BinarySearchNonRecursive(v, 0, len(v)-1, e)
	}
}

func BenchmarkBinarySearch100(b *testing.B) {
	benchmarkBSearch(b, 100)
}

func BenchmarkBinarySearch10000(b *testing.B) {
	benchmarkBSearch(b, 10000)
}

func BenchmarkBinarySearch1000000(b *testing.B) {
	benchmarkBSearch(b, 1000000)
}
