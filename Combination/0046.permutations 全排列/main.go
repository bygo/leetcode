package main

import "sort"

// https://leetcode-cn.com/problems/permutations

func permute(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	if n == 0 {
		return res
	} else if n == 1 {
		return append(res, nums)
	}
	res = [][]int{
		{nums[0], nums[1]},
		{nums[1], nums[0]},
	}

	for k := 2; k < n; k++ {
		num := nums[k]
		resLen := len(res)
		for i := 0; i < resLen; i++ {
			for j := 0; j <= len(res[i]); j++ {
				tmp := []int{}
				if j == 0 {
					tmp = append(append(tmp, num), res[i]...)
				} else {
					t := make([]int, len(res[i][:j]))
					copy(t, res[i][:j])
					tmp = append(append(t, num), res[i][j:]...)
				}
				res = append(res, tmp)
			}
		}
		res = res[resLen:]
	}

	return res
}

func permute(raw []int) [][]int {
	var res [][]int
	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })
	l := len(raw)
	perm := make([]int, l)
	vis := make([]bool, l)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			tmp := make([]int, l)
			copy(tmp, perm)
			res = append(res, tmp)
			return
		}
		for j := 0; j < l; j++ {
			// 已使用 或者 前置未使用，且等于当前值
			if vis[j] {
				continue
			}
			vis[j] = true
			perm[i] = raw[j]
			dfs(i + 1)
			vis[j] = false
		}
	}
	dfs(0)
	return res
}
