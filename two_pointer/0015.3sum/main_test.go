package main

import (
	"testing"
)

var nums = []int{
	-2, 0, 0, 2, 2,
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum(nums)
	}
}