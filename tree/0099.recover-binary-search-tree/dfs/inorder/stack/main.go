package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var last, first, second *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		pre := len(stack) - 1
		if last != nil && stack[pre].Val <= last.Val {
			if first != nil {
				second = stack[pre]
				break //剪枝
			}
			first, second = last, stack[pre]
		}

		last = stack[pre]
		root = stack[pre].Right
		stack = stack[:pre]
	}
	first.Val, second.Val = second.Val, first.Val
}
