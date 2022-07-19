package main

// https://leetcode.cn/problems/merge-in-between-linked-lists

// ❓ 合并两个链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var headLeft = list1
	b = b - a + 2 // 下一个
	for 1 < a {
		list1 = list1.Next
		a--
	}
	tailLeft := list1

	for 0 < b {
		list1 = list1.Next
		b--
	}
	rightHead := list1

	tailLeft.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = rightHead
	return headLeft
}
