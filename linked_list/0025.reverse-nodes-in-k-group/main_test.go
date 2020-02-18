package main

import "testing"

func Test(T *testing.T) {
	l1 := fake([]int{1, 2, 3, 4, 5})
	reverseKGroup(l1, 2)
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
