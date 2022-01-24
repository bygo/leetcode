package main

// https://leetcode-cn.com/problems/missing-number

// ❓ 丢失的数字

func missingNumber(nums []int) int {
	numsL := len(nums)
	numsL = numsL * (numsL + 1) / 2 // 等差数列
	var sum int
	for _, num := range nums {
		sum += num
	}
	return numsL - sum
}
