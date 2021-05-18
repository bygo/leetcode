package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/insertion-sort-list

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	zero := &ListNode{Next: head}
	sorted, cur := head, head.Next
	for cur != nil {
		if sorted.Val <= cur.Val { // 查找待插入值
			sorted = sorted.Next
		} else {
			p := zero
			for p.Next.Val <= cur.Val { // 查找插入点
				p = p.Next
			}
			// 插入
			sorted.Next = cur.Next
			cur.Next = p.Next
			p.Next = cur
		}
		cur = sorted.Next
	}
	return zero.Next
}
