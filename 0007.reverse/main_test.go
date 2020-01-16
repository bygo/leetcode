package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(321)
	}
}

var inputs = map[int]int{
	123:  321,
	-123: -321,
	120:  21,
}

func Test(t *testing.T) {
	for k, v := range inputs {
		if reverse(k) != v {
			t.Error("Error")
		}
	}
}
