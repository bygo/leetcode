package main

// https://leetcode-cn.com/problems/most-frequent-subtree-sum/

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
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sumLeft := dfs(root.Left)
		sumRight := dfs(root.Right)
		sum := sumLeft + sumRight + root.Val
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
