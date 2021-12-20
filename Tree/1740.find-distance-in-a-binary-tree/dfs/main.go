package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	var res int
	c := common(root, p, q)
	path(c, p, q, 0, &res)
	return res
}

func common(root *TreeNode, p, q int) *TreeNode {
	if root == nil || root.Val == p || root.Val == q {
		return root
	}

	l := common(root.Left, p, q)
	r := common(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	return root
}

func path(root *TreeNode, p, q, level int, res *int) {
	if root != nil {
		if root.Val == p || root.Val == q {
			*res += level
		}

		path(root.Left, p, q, level+1, res)
		path(root.Right, p, q, level+1, res)
	}
}
