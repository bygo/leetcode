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
	mid, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = mid
		mid = mid.Next
		fast = fast.Next.Next
	}

	node := &TreeNode{
		Val: mid.Val,
	}

	if mid == head {
		return node
	}

	if pre != nil {
		pre.Next = nil
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node
}
