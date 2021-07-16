package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	zero := &ListNode{} // 前置节点
	prev := zero        // 哨兵节点
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val { // 哪个小 用哪个
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next // 哨兵移动一下
	}

	if l1 == nil {  //
		prev.Next = l2
	} else if l2 == nil {
		prev.Next = l1
	}
	return zero.Next
}
