/*
Package profile demonstrates profiling in Go

  - Benchmarks helps to measure the performance, but where to begin to optimize the performance.
  - Premature optimization is the root of all evil often causing more time in debugging and
    maintainance
  - Programmers waste most of the time thinking about the speed of non critical part of their
    applications. We can forgot about small efficiency improvements 97% of time but optimize for
    the remaining 3% of the time in critical paths.
  - How do we identify the critical paths? Intuitive guesses are mostly wrong when using the
    measurement tools. Profiling identifies the critical parts for improvements.
  - Profiling works based on number of profile events and resulting a statistical summary called
    profile.
  - go test tool has built-in support for several kinds of profiling. It take a particular stack
    trace and measure in different aspects such as CPU,

CPU profile
  - One profile event is recorded for every scheduler interuption
  - Identifies the functions that requires most CPU time. OS scheduler interrupts the thread
  - periodically to schedule the next thead. Each interuption is recorded as one profile event.

Heap profile:
  - One per 512KB of memory allocation
  - Identifies the statements allocating most of the heap memory. profiling library calls the
    internal memory allocation routines to track the allocations.

Blocking profile:
  - One per every gourtine blocking operation
  - Identifies the statements blocking the goroutine most time. This includes system calls, channel
    operations and mutex locks.

Profiling Go Tests

  - Profiling is easy for test code by enabling the flag. Use one flag at a time. Otherwise one
    profiling library may distrub the other results.

  - Applying benchmark functions for profiling can provide better results than a simple tests.

  - Profiling can also be enabled for non-test programs using runtime APIs.

    go test --cpuprofile=cpu.out
    go test --memprofile=mem.out
    go test --blockprofile=block.out

  - The logs generated by profile contains only the addresses of package members. So we need the
    executable to match with the statements. Go test generates an executable named "<package>.test"

Analyzing the profile with pprof:

  - pprof tool can be used to analyze the profile

    go test --run=NONE --bench=BenchmarkBinarySearch --cpuprofile=cpu.log
    go tool pprof --text --nodecount=10 ./profile.test ./cpu.log

  - The --text indicates the output format and --nodecount limits the results to top 10 rows.

  - To generate web based output which requires Graphviz from http://www.graphviz.org/

    go tool pprof --web .\benchmark.test.exe .\cpu.out

Read more about Go profiling here https://go.dev/blog/pprof
*/
package profile

// BinarySearch is a non recursive version of binary search
func BinarySearch(v []int, low int, high int, x int) int {
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
