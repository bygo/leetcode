package main

// https://leetcode-cn.com/problems/increasing-triplet-subsequence

// ❓ 递增三元子序列
// ⚠️ 求 前置第一个最小值 与前置第二个最小值

func increasingTriplet(nums []int) bool {
	numsL := len(nums)
	if numsL < 3 {
		return false
	}

	first, second := nums[0], 1<<31-1
	for i := 1; i < numsL; i++ {
		num := nums[i]
		if second < num {
			// 比第二个大 构成 三元组
			return true
		} else if first < num {
			// 比第一个大 构成 二元组
			second = num
		} else {
			// 比第零个大 构成 一元组
			first = num
		}
	}
	return false
}
