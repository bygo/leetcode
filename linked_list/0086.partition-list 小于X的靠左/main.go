package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/partition-list

func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	leftZero, rightZero := left, right
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}

	right.Next = nil
	left.Next = rightZero.Next
	return leftZero.Next
}
