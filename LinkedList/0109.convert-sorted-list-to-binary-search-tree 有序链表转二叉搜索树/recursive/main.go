package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil { // 找 root
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{
		Val: slow.Val, // slow 为根
	}

	if slow == head { // 只有一个元素
		return root
	}

	if pre != nil { // 与前置断开
		pre.Next = nil
	}

	root.Left = sortedListToBST(head)       // 前置为左
	root.Right = sortedListToBST(slow.Next) // 后置为右
	return root
}
