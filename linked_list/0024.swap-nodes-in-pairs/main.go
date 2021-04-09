package main

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Swap Nodes in Pairs
//Link: https://leetcode-cn.com/problems/swap-nodes-in-pairs

func swapPairs(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	pre := zero
	for head != nil && head.Next != nil {
		left := head
		right := head.Next

		pre.Next = right
		left.Next = right.Next
		right.Next = left

		pre = left
		head = pre.Next
	}

	return zero.Next
}

/**
思路:三指针 zero,left,right
*/
