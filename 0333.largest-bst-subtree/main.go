package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func largestBSTSubtree(root *TreeNode) int {
	res = 0
	dfs(root, -1<<63, 1<<63-1)
	return res
}

func valid(root *TreeNode) {
	if root != nil {
		valid(root.Left)
		cur := dfs(root, -1<<63, 1<<63-1)
		if res < cur {
			res = cur
		}
		valid(root.Right)
	}
}

func dfs(root *TreeNode, min, max int) int {
	if root == nil {
		return 0
	}
	if root.Val < min || max < root.Val {
		return -1
	}
	l := dfs(root.Left, min, root.Val)
	r := dfs(root.Right, root.Val, max)
	if l != -1 && r != -1 {
		cur := l + r + 1
		if res < cur {
			res = cur
		}
		return cur
	}
	return -1
}
