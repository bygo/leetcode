package main

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	var pre [101]int
	res := make([]int, len(nums))
	for i := range nums {
		pre[nums[i]]++
	}

	for i := 1; i < 101; i++ {
		pre[i] += pre[i-1]
	}
	for i, num := range nums {
		if num != 0 {
			res[i] = pre[num-1]
		}
	}
	return res
}
