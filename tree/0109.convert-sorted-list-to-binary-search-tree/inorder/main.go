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
	var l int
	head = h
	for h != nil {
		l++
		h = h.Next
	}
	return convertListToBST(0, l-1)
}

func convertListToBST(l, r int) *TreeNode {
	if r < l {
		return nil
	}
	mid := (l + r) / 2
	node := &TreeNode{
		Val:  head.Val,
		Left: convertListToBST(l, mid-1),
	}
	head = head.Next
	node.Right = convertListToBST(mid+1, r)
	return node
}
