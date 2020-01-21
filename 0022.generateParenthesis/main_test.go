package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("%+v", generateParenthesis(2))
	}
}
