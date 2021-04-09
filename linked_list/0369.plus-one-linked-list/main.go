package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Plus One Linked List
//Link: https://leetcode-cn.com/problems/plus-one-linked-list

func plusOne(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	not := pre
	for head != nil {
		if head.Val != 9 {
			not = head
		}
		head = head.Next
	}

	not.Val += 1
	for not.Next != nil {
		not.Next.Val = 0
		not = not.Next
	}

	if pre.Val == 0 {
		return pre.Next
	} else {
		return pre
	}
}
