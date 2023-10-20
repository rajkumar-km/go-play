/*
Package main demonstrates the concurrency features in Go programming
*/
package main

import "fmt"

// main demonstrates the concurrency features in Go
// - Goroutines
// - WaitGroup
// - Channels
// - Select
func main() {
	fmt.Println("--- DemoGoroutines -------")
  DemoGoroutines()
  fmt.Println("--- DemoWaitGroup -------")
	DemoWaitGroup()
  fmt.Println("--- DemoDemoChannels -------")
	DemoChannels()
  fmt.Println("--- DemoSelect -------")
	DemoSelect()
}
