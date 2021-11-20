package main

// https://leetcode-cn.com/problems/single-number

func singleNumber(nums []int) int {
	var cnt int
	for i := range nums {
		cnt ^= nums[i]
	}
	return cnt
}
