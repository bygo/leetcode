package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Intersection of Two Linked Lists
//Link: https://leetcode-cn.com/problems/intersection-of-two-linked-lists

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}