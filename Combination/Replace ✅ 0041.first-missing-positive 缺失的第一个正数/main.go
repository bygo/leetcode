package main

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	l1 := len(nums)
	for i := 0; i < l1; i++ {
		for 0 < nums[i] && nums[i] <= l1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < l1; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l1 + 1
}
