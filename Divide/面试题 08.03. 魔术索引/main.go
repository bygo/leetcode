package main

// https://leetcode.cn/problems/magic-index-lcci/

func findMagicIndex(nums []int) int {
	var dfs func(lo, hi int) int
	dfs = func(lo, hi int) int {
		if hi < lo {
			return -1
		}
		mid := int(uint(lo+hi) >> 1)
		left := dfs(lo, mid-1)
		if left != -1 {
			return left
		} else if nums[mid] == mid {
			return mid
		}
		return dfs(mid+1, hi)
	}
	return dfs(0, len(nums)-1)
}
