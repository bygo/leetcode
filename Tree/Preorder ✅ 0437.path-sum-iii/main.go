package main

// https://leetcode-cn.com/problems/path-sum-iii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 路径总和 III
func pathSum(root *TreeNode, sumTarget int) int {
	var cnt int
	var sumMpCnt = map[int]int{0: 1}
	var dfs func(node *TreeNode, sumCur int)
	dfs = func(node *TreeNode, sumCur int) {
		if node == nil {
			return
		}
		sumCur += node.Val
		cnt += sumMpCnt[sumCur-sumTarget]
		sumMpCnt[sumCur]++
		dfs(node.Left, sumCur)
		dfs(node.Right, sumCur)
		sumMpCnt[sumCur]-- // 子树遍历完，当前节点的和，不是其他路径的答案
	}
	dfs(root, 0)
	return cnt
}
