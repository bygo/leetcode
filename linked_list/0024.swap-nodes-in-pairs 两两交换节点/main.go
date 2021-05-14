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
		left := head
		right := head.Next

		prev.Next = right
		left.Next = right.Next
		right.Next = left

		prev = left
		head = prev.Next
	}
	return zero.Next
}
