package main

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的前序遍历

func preorder(root *Node) []int {
	var nums []int
	if root == nil {
		return nums
	}
	var stack = []*Node{root}
	for 0 < len(stack) {
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
		nums = append(nums, root.Val)                        //前序输出
		for idx := len(root.Children) - 1; 0 <= idx; idx-- { // 反向
			stack = append(stack, root.Children[idx]) // 入栈
		}
	}
	return nums
}
