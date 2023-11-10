## Go concurrency

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
  - Concurrency is directly built on Go's runtime scheduler - a piece of software.

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

### Race conditions

Let's consider the operations x and y that are executed in different goroutines.
We say it is concurrent if we can not determine which happens first. It could be either way.

A function is called concurrently safe if it continues to work properly even when called by
multiple goroutines.

A type is concurrently safe if all of its methods/operations are concurrently safe.

In a general rule, only concurrently safe types can be accessed by multiple goroutines.

Exported package level functions are generally safe because they have a separate namespace and
don't share the common variables.

Data race happens when multiple goroutines access the same variable and at least one of the
accesses is write.