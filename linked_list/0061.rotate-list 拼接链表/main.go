package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Link: https://leetcode-cn.com/problems/rotate-list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	//tail
	var n = 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	k = n - k%n

	tail.Next = head // 环

	for 1 < k {
		head = head.Next
		k--
	}

	res := head.Next
	head.Next = nil // 断
	return res
}
