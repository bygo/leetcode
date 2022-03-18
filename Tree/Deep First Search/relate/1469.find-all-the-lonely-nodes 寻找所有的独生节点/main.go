package main

// https://leetcode-cn.com/problems/find-all-the-lonely-nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 寻找所有的独生节点

func getLonelyNodes(root *TreeNode) []int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil && node.Right != nil {
			dfs(node.Left)
			dfs(node.Right)
		} else if node.Left != nil {
			nums = append(nums, node.Left.Val)
			dfs(node.Left)
		} else if node.Right != nil {
			nums = append(nums, node.Right.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}
