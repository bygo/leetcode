package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func rightSideView(root *TreeNode) []int {
	res = []int{}
	dfs(root, 1)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) < level {
			res = append(res, root.Val)
		}

		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
}
