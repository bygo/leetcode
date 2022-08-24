package main

// https://leetcode.cn/problems/single-number

func singleNumber(nums []int) int {
	var numUnique int
	for _, num := range nums {
		numUnique ^= num
	}
	return numUnique
}
