package main

// https://leetcode.cn/problems/most-frequent-subtree-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 出现次数最多的子树元素和

func findFrequentTreeSum(root *TreeNode) []int {
	var sumsMax []int
	var cntMax int
	var sumMpCnt = map[int]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		sum := sumLeft + sumRight + node.Val
		sumMpCnt[sum]++
		return sum
	}
	dfs(root)
	for sum, cnt := range sumMpCnt {
		if cnt == cntMax {
			sumsMax = append(sumsMax, sum)
		} else if cntMax < cnt {
			sumsMax = []int{sum}
			cntMax = cnt
		}
	}
	return sumsMax
}
