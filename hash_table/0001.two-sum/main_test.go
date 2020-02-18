package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 26
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
