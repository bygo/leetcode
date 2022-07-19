package main

// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的后序遍历

func postorder(root *Node) []int {
	var nums []int
	if root == nil {
		return nums
	}
	var stack = []*Node{root}
	for 0 < len(stack) {
		nums = append(nums, root.Val)
		for _, n := range root.Children {
			stack = append(stack, n) //入栈
		}

		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}

	lo, hi := 0, len(nums)-1

	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
	return nums
}
