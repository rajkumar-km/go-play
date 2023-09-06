package main

import (
	"testing"
)

// go test . -bench=. -count=5
// PopCountV0 = less than 0.15 ns/op
// PopCountV1 = took 3.7 ns/op
// PopCountV2 = took 17.7 ns/op

func BenchmarkPopCountV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountV0(uint64(i))
	}
}

func BenchmarkPopCountV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountV1(uint64(i))
	}
}

func BenchmarkPopCountV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountV2(uint64(i))
	}
}
