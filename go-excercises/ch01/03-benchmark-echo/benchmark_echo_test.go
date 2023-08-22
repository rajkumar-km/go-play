package main

import "testing"

// go test -bench . | grep -e ns -e Benchmark

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo4()
	}
}
