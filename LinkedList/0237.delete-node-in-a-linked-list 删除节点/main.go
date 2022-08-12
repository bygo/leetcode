package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/delete-node-in-a-linked-list

func deleteNode(node *ListNode) {
	//*node = *node.Next
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
