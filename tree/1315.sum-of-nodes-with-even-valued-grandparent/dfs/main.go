package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func sumEvenGrandparent(root *TreeNode) int {
	sum = 0
	dfs(root, false, false)
	return sum
}

func dfs(root *TreeNode, f bool, g bool) {
	if root != nil {
		if g {
			sum += root.Val
		}
		dfs(root.Left, root.Val%2 == 0, f)
		dfs(root.Right, root.Val%2 == 0, f)
	}
}
