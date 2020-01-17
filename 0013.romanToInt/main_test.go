package main

import (
	"testing"
)

var inputs = map[int]bool{
	123:   false,
	12321: true,
	-123:  false,
}

var s = "LVIII"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt(s)
	}
}

func Test(t *testing.T) {
	h := romanToInt(s)
	println(h)
}
