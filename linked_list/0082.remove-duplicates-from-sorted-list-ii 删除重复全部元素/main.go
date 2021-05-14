package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii

func deleteDuplicates(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	cur := zero
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 出现相同
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v { // 所有相同的都删除
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return zero.Next
}
