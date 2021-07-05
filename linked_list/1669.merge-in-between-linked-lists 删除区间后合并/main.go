package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Link: https://leetcode-cn.com/problems/merge-in-between-linked-lists

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var res = list1
	b -= a + 2 // 从0 开始，并且是下一个
	for 1 < a {
		list1 = list1.Next
		a--
	}
	l := list1

	for 0 < b {
		list1 = list1.Next
		b--
	}
	r := list1

	l.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = r
	return res
}
