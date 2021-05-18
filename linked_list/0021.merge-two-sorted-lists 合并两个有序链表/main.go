package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/merge-two-sorted-lists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	zero := &ListNode{}
	prev := zero
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else if l2 == nil {
		prev.Next = l1
	}
	return zero.Next
}
