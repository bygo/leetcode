package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func sumNumbers(root *TreeNode) int {
	res = 0
	if root != nil {
		dfs(root, 0)
	}
	return res
}

func dfs(root *TreeNode, sum int) {
	if root.Left == nil && root.Right == nil {
		res += sum*10 + root.Val //叶节点 做加法
		return
	}
	if root.Left != nil {
		dfs(root.Left, sum*10+root.Val) //子节点 乘以10
	}
	if root.Right != nil {
		dfs(root.Right, sum*10+root.Val) //子节点 乘以10
	}
}
