/*
rwmutex demonstrates when to use multiple readers, single writer lock.

The sync.Mutex is often used for exclusive locks during reads and writes.
Let's say we have so many read operations, but only have few writes. Since the exclusive
locks are acquired by read operations, writer are slower. In such case, the sync.RWMutex
is useful. This has the method RLock() which can used to allow multiple read operations at
the same without using exclusive locks. This also has the usual Lock() method intended to
be used for write operations with exclusive locks.

Ensure that RLock() is used only for read-only blocks and no writes involved. Because even
some read-only like functions might update something like a counter in shared variable. If
in doubt use exclusive locks

RWMutex involves more bookkeeping and slower than normal Mutex. So it is only profitable to
use with more readers than writers.
*/
package main

import (
	"fmt"
	"sync"
)

var (
	// This is a common pattern to have mutex and the actual variable that it is guarding
	// If you violate this principle, make sure to document it.
	balanceMu sync.RWMutex
	balance   int
)

func Deposit(amount int) {
	balanceMu.Lock()
	// Having the Unlock() in the defer reduces the errors
	// It is not just limited to that. This also makes sure that Unlock() happens even if
	// any panic() occurs. So, use defer for unlock wherever possible
	defer balanceMu.Unlock()
	balance += amount
}

func Withdraw(amount int) bool {
	balanceMu.Lock()
	defer balanceMu.Unlock()
	if balance < amount {
		return false
	}
	balance -= amount
	return true
}

func Balance() int {
	balanceMu.RLock()
	defer balanceMu.RUnlock()
	// Safe to return balance as the deferred RUnlock happens only after fetching the balance
	return balance
}

func main() {
	var wg sync.WaitGroup
	// Perform concurrent deposit and withdrawal operations
	for _, v := range []int{100, 200, -100, 200, -400} {
		wg.Add(1)
		go func(amount int) {
			if amount > 0 {
				Deposit(amount)
				fmt.Println("Deposited amount:", amount)
			} else {
				status := Withdraw(-amount)
				fmt.Println("Withdraw amount:", amount, ":", status)
			}

			defer wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("Final Balance:", Balance())
}
