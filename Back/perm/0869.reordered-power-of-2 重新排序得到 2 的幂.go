package main

import (
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/reordered-power-of-2

// ❓ 重新排序得到 2 的幂

func isPowerOfTwo(num int) bool {
	return num&(num-1) == 0
}

func reorderedPowerOf2(n int) bool {
	var nums []int
	for _, num := range strconv.Itoa(n) {
		nums = append(nums, int(num-'0'))
	}
	sort.Ints(nums)
	numsL := len(nums)
	var vis = make([]bool, numsL)
	var dfs func(idx, num int) bool
	dfs = func(idx, num int) bool {
		if idx == numsL {
			return isPowerOfTwo(num)
		}

		for i := range nums {
			if vis[i] {
				continue
			}

			if num == 0 && nums[i] == 0 {
				continue
			}

			if 0 < i && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}
			vis[i] = true
			if dfs(idx+1, num*10+nums[i]) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return dfs(0, 0)
}
