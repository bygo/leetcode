package main

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	l1 := len(nums)
	// 剔除
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = l1 + 1
		}
	}

	// 记录
	nums = append(nums, l1+1)
	for _, num := range nums {
		num = abs(num)

		if num <= l1 {
			nums[num] = - abs(nums[num])
		}
	}

	// 筛选
	for i := 1; i <= l1; i++ {
		if 0 < nums[i] {
			return i
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
