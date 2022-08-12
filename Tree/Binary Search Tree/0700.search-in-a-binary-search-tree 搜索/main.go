package main

// https://leetcode.cn/problems/search-in-a-binary-search-tree/

// ❓ 二叉搜索树中的搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
