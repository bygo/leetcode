package main

import "fmt"

// https://leetcode.cn/problems/find-duplicate-subtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 寻找重复的子树
// 📚 字符串匹配

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var nodeDup []*TreeNode
	hashMp := map[string][]*TreeNode{}
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		hashL := dfs(node.Left)
		hashR := dfs(node.Right)
		hash := fmt.Sprintf("%d(%s)(%s)", node.Val, hashL, hashR)
		hashMp[hash] = append(hashMp[hash], node)
		return hash
	}
	dfs(root)
	for _, nodes := range hashMp {
		if 2 <= len(nodes) {
			nodeDup = append(nodeDup, nodes[0])
		}
	}
	return nodeDup
}
