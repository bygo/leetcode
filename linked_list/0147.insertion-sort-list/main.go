package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Insertion Sort List
//Link: https://leetcode-cn.com/problems/insertion-sort-list

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ghost := &ListNode{Next: head}
	last, cur := head, head.Next
	for cur != nil {
		if last.Val <= cur.Val {
			last = cur
		} else {
			node := ghost
			for node.Next.Val <= cur.Val {
				node = node.Next
			}
			last.Next = cur.Next
			cur.Next = node.Next
			node.Next = cur
		}
		cur = last.Next
	}
	return ghost.Next
}
