package main

import (
	"testing"
)

var inputs = map[string]string{
	"s":      "s",
	"ssssss": "as*",
	"sssss":  ".*",
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, v := range inputs {
			isMatch(k, v)
		}
	}
}

func Test(t *testing.T) {
	for k, v := range inputs {
		println(isMatch(k, v))
	}
}
