package main

import (
	"testing"
)

var inputs = map[int]string{
	5:   "V",
	3:   "III",
	4:   "IV",
	321: "CCCXXI",
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(123)
	}
}

func Test(t *testing.T) {
	for k, v := range inputs {
		if intToRoman(k) != v {
			t.Error("Error")
		}
	}
}
