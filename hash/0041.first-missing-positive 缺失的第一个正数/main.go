package main

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	l1 := len(nums)
	for i := 0; i < l1; i++ {
		if nums[i] <= 0 {
			nums[i] = l1 + 1
		}
	}

	for i := 0; i < l1; i++ {
		num := abs(nums[i])
		if num <= l1 {
			nums[num-1] = - abs(nums[num-1])
		}
	}

	for i := 0; i < l1; i++ {
		if 0 < nums[i] {
			return i + 1
		}
	}

	return l1 + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
