/*
clock serves the time every second through a TCP channel
It can produce the time based on the timezone set by environment variable CLOCK_TZ
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// bindAddr is the server bind address to listen for clients
const bindAddr = "0.0.0.0"

// loc indicates the clock timezone
var loc *time.Location

func init() {
	// Process the timezone from environment variable
	tz, ok := os.LookupEnv("CLOCK_TZ")
	if !ok {
		// Use the default as UTC
		tz = "UTC"
	}

	var err error
	loc, err = time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Read the port number from command line arguments
	port := flag.Int("port", 8001, "port number")
	flag.Parse()

	// Start the server to listen on particular address
	address := fmt.Sprintf("%s:%d", bindAddr, *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listening on %s (CLOCK_TZ=%s)\n", address, loc.String())

	// Accept the client connections and serve on separate goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

// handleConn serves the time for a single client as long as the connection is alive
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)	
	}
}