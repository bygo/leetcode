package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	if root == nil {
		return res
	}
	dfs(root, nil)
	return res
}

func dfs(root *TreeNode, path []string) {
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(path, "->"))
		return
	}
	if root.Left != nil {
		dfs(root.Left, path)
	}
	if root.Right != nil {
		dfs(root.Right, path)
	}
}
