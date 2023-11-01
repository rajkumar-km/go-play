/*
clockwall displays multiple clocks for different timezones
It is a client that can connect to multiple clock servers serving different timezones

Example:

	clockwall India=localhost:8001 SanJose=localhost:8002 Bristol=localhost:8003
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var clocks []string
	var readers []*bufio.Reader

	if len(os.Args) <= 1 {
		fmt.Println(`Usage  : clockwall Name1=Addr1 Name2=Addr2 ...`)
		fmt.Println(`Example: clockwall India=localhost:8001 SanJose=localhost:8002`)
		os.Exit(1)
	}

	for _, v := range os.Args[1:] {
		tokens := strings.Split(v, "=")
		if len(tokens) != 2 {
			log.Printf("invalid argument: %s", v)
			continue
		}

		address := tokens[1]
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("failed to dial the clock server: %s", address)
			continue
		}

		clocks = append(clocks, tokens[0])
		readers = append(readers, bufio.NewReader(conn))
	}

	if len(clocks) > 0 {
		wallClocks(clocks, readers)
	}
}

// wallClocks displays a table of clocks with different timezones
func wallClocks(clocks []string, readers []*bufio.Reader) {
	counter := 0
	for {
		if (counter % 10) == 0 {
			// Print the header after every 10 seconds
			printHeader(clocks)
		}
		counter++

		for _, r := range readers {
			time, err := r.ReadString('\n')
			if err != nil {
				// Simply display NA when a server is unreachable
				fmt.Printf("%-15s", "NA")
			} else {
				fmt.Printf("%-15s", time[:len(time)-1])
			}
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}

// printHeader displays the header with clock names
func printHeader(clocks []string) {
	for _, name := range clocks {
		fmt.Printf("%-15s", name)
	}
	fmt.Println()
}