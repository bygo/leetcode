package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

func numSubarraysWithSum(nums []int, goal int) int {
	cnt := map[int]int{}
	var sum, res int
	for _, num := range nums {
		cnt[sum]++
		sum += num
		res += cnt[sum-goal]
	}

	return res
}
