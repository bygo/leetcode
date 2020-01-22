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

func swapPairs(head *ListNode) *ListNode {
	empty := &ListNode{}
	empty.Next = head
	prev := empty

	for head != nil && head.Next != nil {
		one := head
		two := head.Next

		prev.Next = two
		one.Next = two.Next
		two.Next = one

		prev = one
		head = one.Next
	}
	return empty.Next
}

/**
思路:三指针 prev,one,two
 */
