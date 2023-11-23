/*
md5sum command produces the md5 digest for the given input string

We need write the command implementation in a separate function so that it is easy for us to test.
  - md5sum() is built as separate function and used in md5sum_test.go
  - Also, using the io.Writer interface variable "out" instead of directly printing on os.Stdout.
    This is required for the test to supply the alternative write buffer and validate the output.
  - Ensure that the code does not call "log.Fatal" or "os.Exit" which can cause the test to abort.
  - Panics are however recovered by the Go test driver.
*/
package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	text = flag.String("t", "", "input text")
)

var md5Out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()
	md5sum(*text)
}

func md5sum(text string) {
	sum := md5.Sum([]byte(text))
	fmt.Fprintf(md5Out, "%x", sum)
}