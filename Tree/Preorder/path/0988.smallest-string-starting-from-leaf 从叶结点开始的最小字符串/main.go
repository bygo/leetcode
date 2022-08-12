package main

import (
	"sort"
)

// https://leetcode.cn/problems/smallest-string-starting-from-leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从叶结点开始的最小字符串

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var strs []string
	var buf []byte
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		buf = append(buf, byte(node.Val+'a'))
		bTop := len(buf) - 1
		if node.Left == nil && node.Right == nil {
			bufCur := []byte{}
			for i := bTop; 0 <= i; i-- {
				bufCur = append(bufCur, buf[i])
			}
			strs = append(strs, string(bufCur))
		}
		dfs(node.Left)
		dfs(node.Right)
		buf = buf[:bTop]
	}
	dfs(root)
	sort.Strings(strs)
	return strs[0]
}
