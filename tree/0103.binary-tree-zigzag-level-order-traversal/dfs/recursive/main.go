package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		if level%2 == 0 {
			dfs(root.Left, level+1)
			dfs(root.Right, level+1)
			//res[level] = append(res[level], root.Val)
		} else {
			dfs(root.Right, level+1)
			dfs(root.Left, level+1)
			//res[level] = append([]int{root.Val}, res[level]...)
		}
	}
}
