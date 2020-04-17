package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int
var sum int

func deepestLeavesSum(root *TreeNode) int {
	max, sum = 0, 0
	dfs(root, 0)
	return sum
}
func dfs(root *TreeNode, level int) {
	if root != nil {
		if max < level {
			max = level
			sum = root.Val
		} else if max == level {
			sum += root.Val
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}
