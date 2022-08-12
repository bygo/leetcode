package main

// https://leetcode.cn/problems/construct-string-from-binary-tree/

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 根据二叉树创建字符串

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == root.Right { // 都是 nil 才会相等
		return strconv.Itoa(root.Val)
	}
	if root.Right == nil { // 右边括号可以省略
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
}
