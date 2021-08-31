package main

// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array

func repeatedNTimes(nums []int) int {
	m := map[int]int{}
	for i := range nums {
		if m[nums[i]] == 1 {
			return nums[i]
		}
		m[nums[i]]++
	}
	return -1
}
