package main

import "fmt"

func main() {
	// Struct embedding
	fmt.Println("--- DemoStructEmbed -------")
	DemoStructEmbed()

	// Embedding pointer to Struct
	fmt.Println("--- DemoPtrToStructEmbed -------")
	DemoPtrToStructEmbed()

	// Embedding interfaces
	fmt.Println("--- DemoInterfaceEmbed -------")
	DemoInterfaceEmbed()
}
