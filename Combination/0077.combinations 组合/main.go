package main

// https://leetcode-cn.com/problems/combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(i int)
	dfs = func(i int) {
		l := len(cur)
		if n-i+l < k {
			return
		}
		if l == k {
			res = append(res, append([]int(nil), cur...))
			return
		}
		cur = append(cur, i+1)
		dfs(i + 1)
		cur = cur[:len(cur)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}

func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(i int)
	dfs = func(i int) {
		l2 := len(cur)

	}
}
