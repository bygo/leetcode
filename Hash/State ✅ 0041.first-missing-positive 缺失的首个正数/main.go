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
		// 原始值当key
		idxNum := abs(num)

		if idxNum <= numsL {
			// 把对应位置上的数字，标记为负数
			// ⚠️ 重复标记，也能保证负数
			nums[idxNum] = - abs(nums[idxNum])
		}
	}

	// 筛选
	for idxNum := 1; idxNum <= numsL; idxNum++ {
		if 0 < nums[idxNum] {
			return idxNum
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
