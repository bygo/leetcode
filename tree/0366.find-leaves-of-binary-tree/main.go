package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func findLeaves(root *TreeNode) [][]int {
	res = nil
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return -1
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	cur := max(l, r) + 1 //拿到当前深度，初叶节点为0，之后不断+1
	if len(res) <= cur {
		res = append(res, []int{})
	}
	res[cur] = append(res[cur], root.Val)
	return cur
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
