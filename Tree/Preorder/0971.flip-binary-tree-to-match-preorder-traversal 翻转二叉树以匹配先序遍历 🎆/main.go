package main

// https://leetcode-cn.com/problems/flip-binary-tree-to-match-preorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 翻转二叉树以匹配先序遍历

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	vL := len(voyage)
	var nums []int
	var dfs func(node *TreeNode)
	var idx int
	dfs = func(node *TreeNode) {
		if node == nil || node.Val != voyage[idx] {
			return
		}
		idx++
		if node.Left != nil && node.Left.Val == voyage[idx] {
			dfs(node.Left)
			dfs(node.Right)
		} else if node.Right != nil && node.Right.Val == voyage[idx] {
			if node.Left != nil {
				// 只有左边存在才需要翻转
				nums = append(nums, node.Val)
			}
			dfs(node.Right)
			dfs(node.Left)
		}
	}
	dfs(root)
	if vL == idx {
		return nums
	}
	return []int{-1}
}
