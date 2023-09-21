/*
Interface composition:
  - Interface can compose the methods of other interfaces
*/
package main

import (
	"bytes"
	"fmt"
)

// A Reader interface provides Read function
type Reader interface {
	Read(p []byte) (int, error)
}

// A Writer interface provides Write function
type Writer interface {
	Write(p []byte) (int, error)
}

// A ReadWriter is a composition of Reader and Writer interfaces
type ReadWriter interface {
	Reader // embed Reader interface
	Writer // embed Writer interface
}

type Disk struct {
	Id int
	buf bytes.Buffer
}
func (d *Disk) Read(p []byte) (int,error) {
	return d.buf.Read(p)	
}
func (d *Disk) Write(p []byte) (int, error) {
	return d.buf.Write(p)
}

func DemoInterfaceEmbed() {
	fmt.Println("Embedding works for interface as well")

	var rwDrive ReadWriter = &Disk{Id: 1}
	n, _ := rwDrive.Write([]byte("hello"))
	fmt.Printf("\tWrite(%q): %d bytes\n", "hello", n)
	
	var p = make([]byte, 128)
	n, _ = rwDrive.Read(p)
	fmt.Printf("\tRead: %d bytes: %s\n", n, string(p))
}