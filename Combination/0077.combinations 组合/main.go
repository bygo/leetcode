package main

// https://leetcode-cn.com/problems/combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(right int)
	dfs = func(right int) {
		l2 := len(cur)
		if l2+n-right+1 < k {
			return
		}
		if l2 == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, right)
		dfs(right + 1)
		cur = cur[:len(cur)-1]
		dfs(right + 1)
	}
	dfs(1)
	return res
}
