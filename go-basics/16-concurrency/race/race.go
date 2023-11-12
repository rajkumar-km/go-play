/*
race demonstrates the Race condition in concurrent programming

- The example involves a global variable "balance" that is shared among routines.
- Multiple deposit transactions are happening at the same using goroutines.
- Finally, we check the balance and it should be sum of all deposits

The program might seems to be correct and it can yield the correct output most of the times.
But there is a possibility of financial loss:

	balance = balance + amount

The above statement involves two concurrent operations: Reading the current balance and
setting the modified balance. What happens when some deposit happens between these two
operations? It is getting lost because we set the balance without accounting the
intermediate transaction.

Many programmers, even some of the clever ones provide justifications for the known data races
in their programs: "the cost of mutual exclusion is high", "it is only used for logging",
"I don't mind losing some messages". But data races are more scenarios it can lead to strange
issues depends on the compiler and platforms. There is no such thing called gentle data race.

How to avoid data races:
1. Do not write to the variable concurrently
  - Concurrent reads do not cause any issues, but having one of the writes bring races.
  - So, avoid writes wherever possible. For instance, avoiding lazy initialization. Instead,
    initialize the entire data structure before starting the goroutines which can only perform
    concurrent reads.

2. Avoid accessing a variable from multiple goroutines
  - "Do not communicate by sharing memory, share memory by communicating"
  - Variables can be confined to a single goroutine and use channels to communicate only
    necessary values to other goroutines.
  - Serial confinement is another method where a single object is sequentially accessed by
    multiple goroutines. For example, a cake object can be accessed first by the baker routine
    and passed to icer routine next. At a time, it can be accessed by only one goroutine.

3. Mutual exclusion using sync.Mutex
  - See mutex/mutex.go for more details
*/
package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) {
	// The following statement involves two concurrent operations: Reading the current balance and
	// setting the modified balance. What happens if the balance is updated between these two
	// transactions. It can result in financial loss if a deposit is made by another goroutine.
	// This is why we need synchronization. See mutex/mutex.go for the fix.
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup

	// Run concurrent deposit transactions
	for _, amount := range []int{100, 200, 300, 400} {
		wg.Add(1)
		go func(v int) {
			Deposit(v)
			fmt.Println("Deposited amount:", v)
			wg.Done()
		}(amount)
	}

	// Finally wait for all the deposits and check the balance
	wg.Wait()
	fmt.Println(Balance())
}
