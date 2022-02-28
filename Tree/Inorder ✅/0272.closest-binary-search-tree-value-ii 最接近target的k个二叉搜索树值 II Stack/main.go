package main

// https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最接近target的k个二叉搜索树值 II

func closestKValues(root *TreeNode, target float64, k int) []int {
	var nums []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := len(stack) - 1
		if len(nums) < k {
			nums = append(nums, stack[top].Val)
		} else if math.Abs(float64(stack[top].Val)-target) < math.Abs(float64(nums[0])-target) {
			nums = append(nums[1:], stack[top].Val)
		} else {
			break
		}
		root = stack[top].Right
		stack = stack[:top]

	}
	return nums
}
