/*
md5sum_test demonstrates testing executable commands in Go

  - Note the test package is also same as main. This is required to modify the md5Out package
    variable to a local buffer so that we can validate the output.
  - go test automatically ignores the main() function defined in the main packages and builds
    its own main to execute the test functions.
*/
package main

import (
	"bytes"
	"testing"
)

func TestMd5sum(t *testing.T) {
	cases := []struct{
		input string
		want string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"abcd", "e2fc714c4727ee9395f324cd2e7f331f"},
	}

	// override md5sum command to write output in our buffer
	buf := new(bytes.Buffer)	
	m5dOutSaved := md5Out
	md5Out = buf
	// restore the original functionality so that its works properly after this test
	defer func() { md5Out = m5dOutSaved }()

	for _, c := range cases {
		buf.Reset()
		md5sum(c.input)
		if buf.String() != c.want {
			t.Errorf("md5sum(%q) = %q, want %q", c.input, buf.String(), c.want)
		}
	}
}
