package main

// https://leetcode-cn.com/problems/find-distance-in-a-binary-tree/

// ❓ 找到二叉树的距离
// ⚠️ 距离 = 边的数量

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	var dist = -1
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil || dist != -1 {
			return -1
		}
		depLeft := dfs(node.Left)
		depRight := dfs(node.Right)
		if depLeft != -1 && depRight != -1 {
			dist = depLeft + depRight + 2
			return -1
		}

		if node.Val == p || node.Val == q {
			if depLeft != -1 {
				dist = depLeft + 1
				return -1
			}
			if depRight != -1 {
				dist = depRight + 1
				return -1
			}
			return 0
		}
		if depLeft != -1 {
			return depLeft + 1
		}
		if depRight != -1 {
			return depRight + 1
		}
		return -1
	}
	dfs(root)
	return dist
}
