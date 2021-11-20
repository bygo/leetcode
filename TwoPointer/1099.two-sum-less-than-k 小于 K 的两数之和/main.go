package main

import "sort"

// https://leetcode-cn.com/problems/two-sum-less-than-k

func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	var max = -1
	for l < r {
		cur := nums[l] + nums[r]
		if cur < k {
			if max < cur {
				max = cur
			}
			l++
		} else {
			r--
		}
	}
	return max
}
