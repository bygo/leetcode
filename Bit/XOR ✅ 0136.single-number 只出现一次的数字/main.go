package main

// https://leetcode-cn.com/problems/single-number

// ❓ 只出现一次的数字

func singleNumber(nums []int) int {
	var num int
	for i := range nums {
		// 两两消除
		num ^= nums[i]
	}
	return num
}
