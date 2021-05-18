package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/swap-nodes-in-pairs

func swapPairs(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	prev := zero
	for head != nil && head.Next != nil {
		l := head
		r := head.Next

		prev.Next = r
		l.Next = r.Next
		r.Next = l

		prev = l
		head = prev.Next
	}
	return zero.Next
}
