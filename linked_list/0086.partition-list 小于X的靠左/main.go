package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/partition-list

func partition(head *ListNode, x int) *ListNode {
	l, r := &ListNode{}, &ListNode{}
	lh, rh := l, r
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			r.Next = head
			r = r.Next
		}
		head = head.Next
	}

	r.Next = nil
	l.Next = rh.Next
	return lh.Next
}
