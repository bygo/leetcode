package main

// https://leetcode.cn/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	numsL := len(nums)
	for idx := 0; idx < numsL; idx++ {
		for 0 < nums[idx] && nums[idx] <= numsL {
			// 正数范围
			// 置换出来 可能也是正数，继续循环
			if nums[nums[idx]-1] == nums[idx] {
				// 对应位置是合法的
				break
			}
			nums[nums[idx]-1], nums[idx] = nums[idx], nums[nums[idx]-1]
		}
	}
	for idx := 0; idx < numsL; idx++ {
		if nums[idx] != idx+1 {
			return idx + 1
		}
	}
	return numsL + 1
}
