package main

import "fmt"

func main() {
	// Struct embedding
	fmt.Println("--- DemoStructEmbed -------")
	DemoStructEmbed()

	// Embedding interfaces
	fmt.Println("--- DemoInterfaceEmbed -------")
	DemoInterfaceEmbed()
}