package main

import (
	"testing"
)

// go test . -bench=. -count=100
// PopCountV1 = less than 0.15 seconds
// PopcountV2 = less than 3.7 seconds !!!

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
