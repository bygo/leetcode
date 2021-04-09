package main

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}


//Title: Merge Two Sorted Lists
//Link: https://leetcode-cn.com/problems/merge-two-sorted-lists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	res := prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	}

	if l2 == nil {
		prev.Next = l1
	}
	return res.Next
}

/**
思路:
根据大小，插入链表
*/
