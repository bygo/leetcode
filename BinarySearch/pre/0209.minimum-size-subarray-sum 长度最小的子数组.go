package pre

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-size-subarray-sum

func minSubArrayLen(target int, nums []int) int {
	numsL := len(nums)
	if numsL == 0 {
		return 0
	}
	subL := 1<<63 - 1
	sums := make([]int, numsL+1)
	for i := 1; i <= numsL; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i <= numsL; i++ {
		idxLeft := sort.SearchInts(sums, target+sums[i-1])
		if idxLeft <= numsL {
			subL = min(subL, idxLeft-i+1)
		}
	}
	if subL == 1<<63-1 {
		return 0
	}
	return subL
}

func minSubArrayLen(target int, nums []int) int {
	numsL := len(nums)
	if numsL == 0 {
		return 0
	}
	subL := 1<<63 - 1
	sums := make([]int, numsL+1)
	for i := 1; i <= numsL; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i <= numsL; i++ {
		idxLeft := sort.SearchInts(sums, target+sums[i-1])
		if idxLeft <= numsL {
			subL = min(subL, idxLeft-i+1)
		}
	}
	if subL == 1<<63-1 {
		return 0
	}
	return subL
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
