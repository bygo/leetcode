package main

// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array

// ❓ 2N 中重复 N 次的元素
// ⚠️ 必定存在[n,n+4] 内有2个重复元素

func repeatedNTimes(nums []int) int {
	numsL := len(nums)
	for step := 1; step <= 3; step++ {
		for i := 0; i < numsL-step; i++ {
			if nums[i] == nums[i+step] {
				return nums[i]
			}
		}
	}

	return -1
}
