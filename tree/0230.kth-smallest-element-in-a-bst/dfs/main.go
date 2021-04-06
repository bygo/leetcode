package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	dfs(root, &k, &res)
	return res
}

func dfs(root *TreeNode, k *int, res *int) {
	if root != nil {
		dfs(root.Left, k, res)
		*k--
		if *k == 0 {
			*res = root.Val
			return
		}
		dfs(root.Right, k, res)
	}
}
