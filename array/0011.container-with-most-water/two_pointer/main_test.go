package main

import (
	"testing"
)

var inputs = map[int][]int{
	49: []int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	},
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			maxArea(input)
		}
	}
}

func Test(t *testing.T) {
	for v, input := range inputs {
		if maxArea(input) != v {
			t.Error("Error")
		}
	}
}
