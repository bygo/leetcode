package main

import (
	"testing"
)

var inputs = map[int]bool{
	123:   false,
	12321: true,
	-123:  false,
}

var head = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeNthFromEnd(head, 1)
	}
}

func Test(t *testing.T) {
	h := removeNthFromEnd(head, 1)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}
