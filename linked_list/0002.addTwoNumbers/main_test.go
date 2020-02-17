package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	l1 := fake([]int{5, 6, 7})
	l2 := fake([]int{5, 5, 5})
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
