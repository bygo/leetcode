package main

import (
	"testing"
)

var nums = []int{
	-1, 2, 1, -4,
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSumClosest(nums, 1)
	}
}

func Test(t *testing.T) {
	r := threeSumClosest(nums, 1)
	println(r)
}
