package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

func numSubarraysWithSum(nums []int, goal int) int {
	var res, sum int
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[sum]++
		sum += num
		res += cnt[sum-goal] // 前缀和 sum[j] - sum[i] = goal
	}
	return res
}
