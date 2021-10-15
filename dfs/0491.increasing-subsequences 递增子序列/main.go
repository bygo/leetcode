package main

import "math"

// https://leetcode-cn.com/problems/increasing-subsequences

var (
	l    int
	tmp  []int
	mod  = int(1e9 + 7)
	base = 77
)

func findSubsequences(nums []int) [][]int {
	l = len(nums)
	n := 1 << l
	res := [][]int{}
	m := map[int]bool{}
	m2 := map[int]bool{}
	for _, num := range nums {
		m2[num] = true
	}
	var h int
	mask := 0
	for i := 0; i < n; i++ {
		mask = i
		tmp = []int{}
		for i := 0; i < l; i++ {
			if (mask & 1) != 0 {
				tmp = append(tmp, nums[i])
			}
			mask >>= 1
		}

		for _, x := range tmp {
			h = h*base%mod + x + base
			h %= mod
		}

		if !m[h] && check() {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			m[h] = true
		}
	}
	return res
}

func check() bool {
	for i := 1; i < len(tmp); i++ {
		if tmp[i] < tmp[i-1] {
			return false
		}
	}
	return 2 <= len(tmp)
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(index, last int)
	var l1 = len(nums)
	dfs = func(index, last int) {
		if l1 == index {
			l2 := len(tmp)
			if 2 <= l2 {
				t := make([]int, l2)
				copy(t, tmp)
				res = append(res, t)
			}
			return
		}

		if last <= nums[index] {
			tmp = append(tmp, nums[index])
			dfs(index+1, nums[index])
			tmp = tmp[:len(tmp)-1]
		}
		if last != nums[index] {
			dfs(index+1, last)
		}
	}
	dfs(0, math.MinInt32)
	return res
}
