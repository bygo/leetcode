package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Middle of the Linked List
//Link: https://leetcode-cn.com/problems/middle-of-the-linked-list

func middleNode(head *ListNode) *ListNode {
	var left, right = head, head

	for right.Next != nil && right.Next.Next != nil {
		left = left.Next
		right = right.Next.Next
	}

	if right.Next != nil {
		return left.Next
	}

	return left
}
