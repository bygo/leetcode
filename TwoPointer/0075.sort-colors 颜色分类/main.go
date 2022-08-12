package main

// https://leetcode.cn/problems/sort-colors/

func sortColors(nums []int) {
	var l int
	var r = len(nums) - 1
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
			i--
		}
	}
}

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		for i <= r && nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}
