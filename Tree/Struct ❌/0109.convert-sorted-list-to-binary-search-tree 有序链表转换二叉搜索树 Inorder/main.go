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

func sortedListToBST(link *ListNode) *TreeNode {
	var linkL int
	head := link
	for link != nil {
		linkL++
		link = link.Next
	}

	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if right < left {
			return nil
		}
		mid := (left + right) / 2
		node := &TreeNode{}
		node.Left = dfs(left, mid-1)
		node.Val = head.Val
		head = head.Next
		node.Right = dfs(mid+1, right)
		return node
	}
	return dfs(0, linkL-1)
}
