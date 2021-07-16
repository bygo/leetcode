package main

// https://leetcode-cn.com/problems/missing-number

func missingNumber(nums []int) int {
	l1 := len(nums)
	l1 = l1 * (l1 + 1) / 2
	var sum int
	for _, num := range nums {
		sum += num
	}
	return l1-sum
}
