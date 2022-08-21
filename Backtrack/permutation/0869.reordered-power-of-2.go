package main

import (
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/reordered-power-of-2

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
	var used = make([]bool, numsL)
	var dfs func(idx, num int) bool
	dfs = func(idx, num int) bool {
		if idx == numsL {
			return isPowerOfTwo(num)
		}

		for i := range nums {
			if used[i] {
				continue
			}

			if num == 0 && nums[i] == 0 {
				continue
			}

			if 0 < i && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			if dfs(idx+1, num*10+nums[i]) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return dfs(0, 0)
}
