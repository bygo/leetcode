package main

// https://leetcode-cn.com/problems/missing-number

// ❓ 丢失的数字
// ⚠️ 0 ~ numsL 丢失了一个数字
func missingNumber(nums []int) int {
	numsL := len(nums)
	var num int

	// 起始
	num ^= nums[0]
	// 中间diff
	for n := 1; n < numsL; n++ {
		num ^= n
		num ^= nums[n]
	}
	// 尾部 numsL
	return num ^ numsL
}
