package main

// https://leetcode-cn.com/problems/first-missing-positive/

// ❓ 缺失的首个正数

func firstMissingPositive(nums []int) int {
	numsL := len(nums)
	// 剔除
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = numsL + 1
		}
	}

	// 记录
	nums = append(nums, numsL+1)
	for _, num := range nums {
		num = abs(num)

		if num <= numsL {
			nums[num] = - abs(nums[num])
		}
	}

	// 筛选
	for i := 1; i <= numsL; i++ {
		if 0 < nums[i] {
			return i
		}
	}
	return numsL + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
