package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res, _ := dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) (*TreeNode, int) {
	if root == nil {
		return root, level
	}
	l, lHeight := dfs(root.Left, level+1)
	r, rHeight := dfs(root.Right, level+1)

	if lHeight == rHeight {
		return root, lHeight
	}

	if lHeight < rHeight {
		return r, rHeight
	}

	return l, lHeight
}
