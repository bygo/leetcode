package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list

func deleteDuplicates(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	cur := zero
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 出现相同
			cur.Next = cur.Next.Next // 去掉前一个
		} else {
			cur = cur.Next
		}
	}
	return zero.Next
}
