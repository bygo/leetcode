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

var head *ListNode

func sortedListToBST(h *ListNode) *TreeNode {
	var s int
	head = h
	for h != nil {
		s++
		h = h.Next
	}
	return convertListToBST(0, s-1)
}

func convertListToBST(l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	left := convertListToBST(l, mid-1)
	node := &TreeNode{
		Val: head.Val,
	}
	node.Left = left
	head = head.Next
	node.Right = convertListToBST(mid+1, r)
	return node
}
