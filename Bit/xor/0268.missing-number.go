package main

// https://leetcode.cn/problems/missing-number

// ❓ 丢失的数字
// ⚠️ 0 ~ numsL 丢失了一个数字
func missingNumber(nums []int) int {
	numsL := len(nums)
	var num int

	for idx := 0; idx < numsL; idx++ {
		num ^= idx
		num ^= nums[idx] // case: 0
	}
	return num ^ numsL
}
