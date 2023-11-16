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

### Mutual Exclusion (sync.Mutex)
- See mutex/mutex.go

### Read/Write Mutexes (sync.RWMutex)
- See rwmutex/rwmutex.go

### Memory Synchronization

Do you wonder why we need locks for read operations?
1. First thing is that a read may give wrong result while the write is ongoing
2. Secondly, the more subtle is flushing out memory for synchronization between goroutines.

In the moderen processors, dozens for CPU has its own cache and defer the writes to main memory.
So, a value modified in one goroutine may not be reflected immediately in another goroutine.
Synchronization calls like channel operations and mutex flush out the cache to memory. This
ensures that the updated values are reflected in other goroutines running on different
processors.

Compiler on the other side can perform optimization and it can even rearrange statements as
long the output is same. So, the intutions about the concurrency are not be trusted.

### Lazy Initialization (sync.Once)
- See once/once.go

### Race Detector
It is possible for the programmers to leak concurrency issues even after the review. Go provides
a built-in race detector to detect the concurrency issues at run time.

- Simple add "-race" flag in go build which would generate a detailed executable
- Race detector analyzes the go events/statements while running the program and detect the
  possible races. It checks whether a shared variable is accessed by multiple goroutines in
  a short span of time without using the synchronization methods.
- Since all the go events are recorded, errors are reported with the information about shared
  variable, and the stack trace of two goroutines accessing the same variable.
- Note that it can detect only races that happens during the run. So, ensure that all the paths
  are covered in the tests and -race flag is used.
- Race detector involves extra bookkeeping and affects performance. But, most of the time it is
  good to run production systems with race detector which is tolerable. This can save lots of
  debugging time.

### Shared variables with locks vs Communicating sequencial processes (channels)

"It’s not always obvious which approach is preferable in a given situation, but it’s worth
knowing how they correspond. Sometimes switching from one approach to the other can
make your code simpler."

### Goroutines and Threads
- Threads are the underlying operating system units to execute a piece of code
- Goroutine is an abstraction on top of OS threads to achieve certain benefits.
- The benifits of quantitative: Basically M goroutines are mapped to N OS threads (called m:n)
- Reducing large number of OS threads becomes a qualitative benefits as well.
#### Growable stacks
- Threads are fixed in size typically 2 MB. All the stack variables are stored in this memory
  This can be too much or too less depends on the requirements.
- Thread size can be configured using ulimit or during the link stage or at run time before
  creating a thread. However the size is fixed once set.
- Goroutines grows and shrinks based on needs starting from 2 KB. It may grow even to 1 GB
- In terms of size, we can create 1000s of smalller goroutines in place of a single OS thread
#### Goroutine scheduling
- OS threads are scheduled by the OS kernel:
    - Every few milliseconds, a hardware timer interupts the processor
    - Processor performs the context switch to next thread.
    - Context switching is costly as it has to suspend the current thread and save all its
      registers in memory, and then reload the next available thread.
- Go runtime has its own Go scheduler that works by m:n scheduling
    - Go scheduler multiplexes/schedules m goroutines to n OS threads.
    - Unlike OS thread scheduler, Go scheduler is not invoked by timers.
    - This is specific to the single program and invoked during certain calls such as
      time.Sleep(), channel read/write, or mutex locks.
    - Go scheduler works like a OS thread scheduler during this time and performs the
      context switch between goroutines.
    - Context switch of goroutine is much cheaper than the full context switch of OS thread.
#### GOMAXPROCS
- Go scheduler uses the parameter GOMAXPROCS to limit the active OS threads to be used.
  Note that this is the limit of active OS threads. It does not count any OS threads blocked by
  user input or system calls.
- User can control this by setting environment variable GOMAXPROCS or using the function
  runtime.GOMAXPROCS
#### Goroutines have no identity
- OS threads generally has a unique identifier.
- Thread local storage can be built using the unique thread id and each thread can use their
  dedicated variable space.
- Goroutine does not any unique identifier. This is intentional to avoid routines building
  individual states like thread-local storage.
- Instead of using state variables, Go encourage to use arguments that is more explicit and easy
  for understanding and debugging.
- All goroutines works straightforward. We don't have to worry about individual goroutine states.