/*
go test .\10-comma\ --bench=. --count=3
*/
package main

import "testing"

func BenchmarkComma(b *testing.B) {
	for i:=0; i < b.N; i++ {
		comma("123456789")
	}
}