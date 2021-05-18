package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历栈 + 滑动窗口
func closestKValues(root *TreeNode, target float64, k int) []int {
	var res []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := len(stack) - 1
		if len(res) < k {
			res = append(res, stack[top].Val)
		} else if math.Abs(target-float64(stack[top].Val)) < math.Abs(target-float64(res[0])) {
			res = append(res[1:], stack[top].Val)
		}
		root = stack[top].Right
		stack = stack[:top]

	}
	return res
}
