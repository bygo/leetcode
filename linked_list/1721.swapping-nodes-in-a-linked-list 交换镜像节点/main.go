package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list

func swapNodes(head *ListNode, k int) *ListNode {
	var slow, fast, left = head, head, head
	for 1 < k {
		fast = fast.Next
		k--
	}

	left = fast

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	left.Val, slow.Val = slow.Val, left.Val

	return head
}
