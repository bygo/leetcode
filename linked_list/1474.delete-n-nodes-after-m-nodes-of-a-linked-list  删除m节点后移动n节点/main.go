package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Link: https://leetcode-cn.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list

// 直接模拟
func deleteNodes(head *ListNode, m int, n int) *ListNode {
	res := head
	for head != nil {
		for i := 1; i < m && head != nil; i++ {
			head = head.Next
		}
		pre := head

		for j := -1; j < n && head != nil; j++ {
			head = head.Next
		}

		if pre != nil {
			pre.Next = head
		}
	}

	return res
}
