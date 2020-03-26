package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	dfs(root, sum, []int{})
	return res
}

func dfs(root *TreeNode, sum int, stack []int) {
	if root == nil {
		return
	}
	stack = append(stack, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			r := make([]int, len(stack))
			copy(r, stack)
			res = append(res, r)
		}
	}
	dfs(root.Left, sum-root.Val, stack)
	dfs(root.Right, sum-root.Val, stack)
}
