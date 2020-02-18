package main

import (
	"testing"
)

var s = []string{
	"flower", "flow", "flight", "f",
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(s)
	}
}
