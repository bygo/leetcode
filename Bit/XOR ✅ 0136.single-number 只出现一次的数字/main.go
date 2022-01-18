package main

// https://leetcode-cn.com/problems/single-number

// ❓ 只出现一次的数字

func singleNumber(nums []int) int {
	var numUnique int
	for _, num := range nums {
		// 两两消除
		numUnique ^= num
	}
	return numUnique
}
