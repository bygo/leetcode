package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/plus-one-linked-list

func plusOne(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	cur := zero
	// 找非9
	for head != nil {
		if head.Val != 9 {
			cur = head
		}
		head = head.Next
	}

	cur.Val += 1

	// cur后 所有 9 转 0
	for cur.Next != nil {
		cur.Next.Val = 0
		cur = cur.Next
	}

	// 是否需要进位
	if zero.Val == 0 {
		return zero.Next
	} else {
		return zero
	}
}
