package main

import "fmt"

func main() {
	// Struct embedding
	fmt.Println("--- DemoStructEmbed -------")
	DemoStructEmbed()

	// Embedding pointer to struct
	fmt.Println("--- DemoPtrToStructEmbed -------")
	DemoPtrToStructEmbed()

	// Embedding in unnamed struct
	fmt.Println("--- DemoUnnamedStructEmbed -------")
	DemoUnnamedStructEmbed()

	// Embedding interfaces
	fmt.Println("--- DemoInterfaceEmbed -------")
	DemoInterfaceEmbed()
}
