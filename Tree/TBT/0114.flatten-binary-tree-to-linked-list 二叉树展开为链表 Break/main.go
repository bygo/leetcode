package main

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树展开为链表
// ⚠️ 与先序遍历相同

func flatten(root *TreeNode) *TreeNode {
	head := root
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right
			root.Left = nil
		}
	}
	return head
}
