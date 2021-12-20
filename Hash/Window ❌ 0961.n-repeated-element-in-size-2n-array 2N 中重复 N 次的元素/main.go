package main

// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array

func repeatedNTimes(nums []int) int {
	l1 := len(nums)
	for j := 1; j <= 2; j++ {
		for i := 0; i < l1-j; i++ {
			if nums[i] == nums[i+j] {
				return nums[i]
			}
		}
	}

	return -1
}
