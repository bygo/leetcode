package main

// https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树中的众数

func findMode(root *TreeNode) []int {
	var cntMax, cnt int
	var numsMode []int
	var numCur = -1 << 63
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if numCur < node.Val {
			numCur = node.Val
			cnt = 1
		} else {
			cnt++
		}
		if cntMax == cnt {
			numsMode = append(numsMode, numCur)
		} else if cntMax < cnt {
			numsMode = []int{numCur}
			cntMax = cnt
		}
		dfs(node.Right)
	}
	dfs(root)
	return numsMode
}
