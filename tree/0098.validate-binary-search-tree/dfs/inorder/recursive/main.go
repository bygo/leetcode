package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
dfs
1.左节点小于根节点
2.右节点大于根节点
3.左节点的右节点小于根节点
（因为 根节点的左节点下的所有值，必须全部小于根节点）
4.右节点的左节点大于根节点
（因为 根节点的右节点下的所有值，必须全部大于根节点）
 */

var last *TreeNode

func isValidBST(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root
	return dfs(root.Right)
}
