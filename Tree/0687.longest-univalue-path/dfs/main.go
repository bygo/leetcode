package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	dfs(root, 0, 0)
	return res
}

func dfs(root *TreeNode, f int, path int) int {
	if root != nil {
		l := dfs(root.Left, root.Val, path)  //左树最长路径
		r := dfs(root.Right, root.Val, path) //右树最长路径
		res = max(r+l, res)                  //路径是否变长
		if f == root.Val {                   //相等时+1，并且选择l或r中最长路径
			return max(l, r) + 1
		}
	}
	return 0 //其他情况全部返回0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
