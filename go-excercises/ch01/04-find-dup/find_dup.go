/*
find_dup reads the lines from given input file(s) or from the standard input,
prints only the duplicate lines.
  - bufio.Scanner is used to read the input line by line.
  - map[string]bool is used to detect the duplicate lines
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, filePath := range os.Args[1:] {
			file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			printDuplicates(file, filePath)
		}
	} else {
		fmt.Println("No input files provided. Reading from standard input.")
		printDuplicates(os.Stdin, "stdin")
	}
}

// printDuplicates prints the duplicate lines in a given file stream
// This also prints the tag as prefix
func printDuplicates(r io.Reader, tag string) {
	var lines map[string]bool = make(map[string]bool)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Error reading", scanner.Err())
			return
		}

		line := scanner.Text()
		printed, ok := lines[line]

		if !ok {
			// Initialize the key for the first time and set printed=false
			lines[line] = false
		} else if !printed {
			// Line already occurred once, print it on the second occurance
			fmt.Printf("%s: %s\n", tag, line)
			lines[line] = true
		}
	}
}
