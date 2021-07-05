package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Merge In Between Linked Lists
// Link: https://leetcode-cn.com/problems/merge-in-between-linked-lists

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var head = list1
	b = b - a + 2
	for 1 < a {
		list1 = list1.Next
		a--
	}
	leftTail := list1

	for 0 < b {
		list1 = list1.Next
		b--
	}
	rightHead := list1

	leftTail.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = rightHead
	return head
}