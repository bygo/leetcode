package main

// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的前序遍历

func preorder(root *Node) []int {
	var nums []int
	var stack = []*Node{}
	for root != nil {
		for root != nil {
			nums = append(nums, root.Val)
			if len(root.Children) == 0 {
				root = nil
			} else {
				for idx := len(root.Children) - 1; 0 < idx; idx-- { // 反向
					stack = append(stack, root.Children[idx]) // 入栈
				}
				root = root.Children[0]
			}
		}

		sTop := len(stack) - 1
		for 0 <= sTop && root == nil {
			root = stack[sTop]
			stack = stack[:sTop]
			sTop--
		}
	}
	return nums
}
