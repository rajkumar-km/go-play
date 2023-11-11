/*
mutex demonstrates using mutually exclusive locks for concurrency
*/
package main

import (
	"fmt"
	"sync"
)

var (
	// This is a common pattern to have mutex and the actual variable that it is guarding
	// If you violate this principle, make sure to document it.
	balanceMu sync.Mutex
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
	balanceMu.Lock()
	defer balanceMu.Unlock()
	// Safe the return balance as the deferred Unlock happens only after fetching the balance
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
