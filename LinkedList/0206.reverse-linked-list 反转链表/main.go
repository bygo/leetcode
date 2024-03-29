package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
