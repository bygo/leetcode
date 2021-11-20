package main

// https://leetcode-cn.com/problems/largest-unique-number

func largestUniqueNumber(nums []int) int {
	m := [1001]int{}
	for i := range nums {
		m[nums[i]]++
	}

	for i := 1000; 0 <= i; i-- {
		if m[i] == 1 {
			return i
		}
	}
	return -1
}
