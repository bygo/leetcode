package main

import (
	"fmt"
	"testing"
)

var nums = []int{
	1, 0, -1, 0, -2, 2,
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum(nums, 0)
	}

}

func Test(t *testing.T) {
	fmt.Printf("%+v", fourSum(nums, 0))
}
