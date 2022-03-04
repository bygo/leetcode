package main

// https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 统计二叉树中好节点的数目
// ⚠️ 根到节点 X 路径 没有比 X 大
func goodNodes(root *TreeNode) int {
	var cntGood int
	var dfs func(node *TreeNode, max int)
	dfs = func(node *TreeNode, max int) {
		if node != nil {
			if max <= node.Val {
				max = node.Val
				cntGood++
			}
			dfs(node.Left, max)
			dfs(node.Right, max)
		}
	}
	dfs(root, root.Val)
	return cntGood
}
