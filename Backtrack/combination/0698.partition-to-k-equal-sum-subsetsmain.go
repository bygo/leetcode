package main

import "sort"

// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%k != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	cnt := total / k

	mL := len(nums)
	edges := make([]int, k)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == mL {
			return true
		}

		for i := range edges {
			if 0 < i && edges[i-1] == edges[i] {
				continue
			}
			edges[i] += nums[idx]
			if edges[i] <= cnt && dfs(idx+1) {
				return true
			}
			edges[i] -= nums[idx]
		}
		return false
	}
	return dfs(0)
}
