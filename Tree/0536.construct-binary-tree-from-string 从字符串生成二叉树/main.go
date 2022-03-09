package main

import "strconv"

// https://leetcode-cn.com/problems/construct-binary-tree-from-string

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从字符串生成二叉树

func str2tree(s string) *TreeNode {
	var dfs func(s string) *TreeNode
	dfs = func(s string) *TreeNode {
		if s == "" {
			return nil
		}
		node := &TreeNode{}
		var idxL, idxR, cnt int
		sL := len(s)
		for i := 0; i < sL; i++ {
			if s[i] == '(' {
				if cnt == 0 {
					idxL = i
				}
				cnt++
			} else if s[i] == ')' {
				cnt--
				if cnt == 0 {
					idxR = i
					break
				}
			}
		}
		if idxL == idxR {
			node.Val, _ = strconv.Atoi(s)
			return node
		}
		node.Val, _ = strconv.Atoi(s[:idxL])
		node.Left = dfs(s[idxL+1 : idxR])
		if idxR+1 < sL {
			node.Right = dfs(s[idxR+2 : sL-1])
		}

		return node
	}
	return dfs(s)
}
