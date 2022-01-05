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
		// 原始值
		num = abs(num)

		if num <= numsL {
			// 存在标记为负数，相反，循环后第一个正数即为答案
			nums[num] = - abs(nums[num])
		}
	}

	// 筛选
	for i := 1; i <= numsL; i++ {
		if 0 < nums[i] {
			return i
		}
	}

	// 1~numsL 都存在(都是负数)
	return numsL + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
