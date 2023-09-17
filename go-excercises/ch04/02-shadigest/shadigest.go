/*
shadigest produces the 256/384/512 bits digest for the given input text

Usage:

	shadigest [sha256|sha384|sha512]

Default: sha256
*/
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

const (
	Sha256 = "sha256"
	Sha384 = "sha384"
	Sha512 = "sha512"
)

func main() {
	// Parse command line arguments to get the digest size
	shaType := Sha256
	if len(os.Args) > 1 {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: shadigest [sha256|sha384|sha512]\n")
			os.Exit(1)
		}
		shaType = os.Args[1]
	}

	// Read the message from standard input
	fmt.Println("Enter the message and press ctrl+D")
	msg, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading lines: %s\n", err.Error())
		os.Exit(1)
	}

	// Get the sha digest for the message
	sum, err := shaDigest(msg, shaType)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	// Note that %x concatenates all the elements of byte slice and prints as single digest
	fmt.Printf("%s = %x\n", shaType, sum)
}

// shaDigest calculates the sha digest for the given msg and shaType.
// Returns error if shaType is other than 256, 384, or 512.
func shaDigest(msg []byte, shaType string) ([]byte, error) {
	switch shaType {
	case Sha256:
		sum := sha256.Sum256(msg)
		return sum[:], nil
	case Sha384:
		sum := sha512.Sum384(msg)
		return sum[:], nil
	case Sha512:
		sum := sha512.Sum512(msg)
		return sum[:], nil
	default:
		return nil, fmt.Errorf("sha type %s is unsupported", shaType)
	}
}
