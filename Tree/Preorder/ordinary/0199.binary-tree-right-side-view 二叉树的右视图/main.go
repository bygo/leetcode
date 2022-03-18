package main

// https://leetcode-cn.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的右视图

func rightSideView(root *TreeNode) []int {
	var nums []int
	var dfs func(node *TreeNode, dep int)
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}

		if len(nums) < dep {
			nums = append(nums, node.Val)
		}

		dfs(node.Right, dep+1)
		dfs(node.Left, dep+1)
	}
	dfs(root, 1)
	return nums
}
