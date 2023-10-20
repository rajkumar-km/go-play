/*
goroutines demonstrates simple goroutine in Go

Concurrency is directly built on Go's runtime scheduler - a piece of software.

Goroutines:

  - Go has goroutines which is lightweight user space threads
  - Lightweight and cheaper than kernel threads
  - Smaller memory footprint. Initial memory is 2KB while kernel thread is 8KB
  - Faster creation, destruction, and context switches
  - Any function invoked as goroutine is considered independent, scheduled
    and executed
  - OS can schedule only OS threads to processors.
  - Go assigns the OS thread to goroutines when scheduled. If a goroutine is
    blocked for I/O, Network, System call then the OS thread is assigned to
    another goroutine. OS thread is again allocated to the goroutine once is
    unblocked
  - So, we can perform concurrent operations even with a single OS thread.
    This is called concurrency.

Concurrency:
  - Performing multiple tasks in the overlapping time periods. Go has
    concurrency since it can run multiple go routines with a single OS thread.

Parallelism:
  - Performing multiple tasks at the same time using multiple system cores

Go scheduler:

 1. Reuse threads

 2. Limit creation of kernel threads to number of CPU cores
    - This limit does not count the blocked kernel threads.

 3. Distributed runqueue with stealing and handoff
    - Distributed runqueue created on per core basis and is stored in a Heap struct.
    - Scheduler can can steal half of other queue when a work queue is empty.
    - Go does not block the kernel thread for channel read/write operations.
    - But, scheduler does not aware of the blocking system calls and this
    can block the kernel thread itself.
    - A background monitor thread sysmon watches a blocked thread and handoff its queue to another thread.

 4. A global runqueue with low priority
    - Some goroutines may run heavy jobs that can bring the entire system down.
    - sysmon a background thread and preemption to detects the long running goroutines (> 10ms)
    - Unschedule them when possible and put in global runqueue

 5. Threads without work look for tasks before parking
    - Look at Global runqueue
    - Run garbage collection tasks
    - Work stealing

Priority based goroutines are not supported like linux scheduler.
Not aware of system topologoy like NUMA
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Main routine starts

	// spinner becomes a new goroutine
	go spinner(100 * time.Millisecond)

	// Main routine runs a job
	res := fib(45)
	fmt.Printf("\r45th fibonacci number is: %d\n", res)

	// Exit of main routine, also kills other goroutines
}

// fib returns the xth fibonacci number
func fib(x uint64) uint64 {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// spinner displays a loading text animation
func spinner(delay time.Duration) {
	for {
		for _,r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}