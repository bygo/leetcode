package main

// https://leetcode-cn.com/problems/binary-tree-paths/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的所有路径

func binaryTreePaths(root *TreeNode) []string {
	var strs []string
	if root == nil {
		return strs
	}
	var paths []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		strs = append(strs, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			paths = append(paths, strings.Join(strs, "->"))
		}
		dfs(node.Left)
		dfs(node.Right)
		strs = strs[:len(strs)-1]
	}
	dfs(root)
	return paths
}
