package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	l1 := fake([]int{5, 6, 7})
	l2 := fake([]int{5, 5, 5})
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

func fake(list []int) *ListNode {
	l := &ListNode{}
	o := l
	for k, v := range list {
		o.Val = v
		if len(list) > k+1 {
			o.Next = &ListNode{}
			o = o.Next
		}
	}
	return l
}
