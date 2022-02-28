package main

// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 有序链表转换二叉搜索树

func sortedListToBST(left *ListNode) *TreeNode {
	if left == nil {
		return nil
	}
	var pre *ListNode
	slow, fast := left, left
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := &TreeNode{
		Val: slow.Val,
	}

	if slow == left {
		return node
	}

	if pre != nil {
		pre.Next = nil
	}

	node.Left = sortedListToBST(left)
	node.Right = sortedListToBST(slow.Next)
	return node
}
