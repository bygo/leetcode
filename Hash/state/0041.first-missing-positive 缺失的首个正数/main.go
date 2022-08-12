package main

// https://leetcode.cn/problems/first-missing-positive/

// ❓ 缺失的首个正数

func firstMissingPositive(nums []int) int {
	numsL := len(nums)
	// 剔除
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = numsL + 1
		}
	}

	// 存在标记
	//nums = append(nums, numsL+1)
	for _, num := range nums {
		// 原始值 偏移(-1) 当key
		idxNum := abs(num) - 1

		if idxNum < numsL {
			// 把对应位置上的数字，标记为负数
			// ⚠️ 重复标记，也能保证负数
			nums[idxNum] = -abs(nums[idxNum])
		}
	}

	// 筛选 数字 1~numsL (偏移-1)
	for num := 0; num < numsL; num++ {
		if 0 < nums[num] {
			return num + 1
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
