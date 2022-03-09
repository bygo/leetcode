package main

// https://leetcode-cn.com/problems/all-possible-full-binary-trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 所有可能的满二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func allPossibleFBT(n int) []*TreeNode {
	var f []*TreeNode
	if n%2 == 0 {
		return f
	}
	if n == 1 {
		return append(f, &TreeNode{})
	}
	for i := 1; i < n-1; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - 1 - i)
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				root := new(TreeNode)
				root.Left = left[j]
				root.Right = right[k]
				f = append(f, root)
			}
		}
	}
	return f
}
