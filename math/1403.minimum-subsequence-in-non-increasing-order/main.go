package main

import "sort"

func minSubsequence(nums []int) []int {
	var res []int
	sort.Ints(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	counter := 0
	i := len(nums)
	half := sum / 2

	for counter <= half {
		i--
		counter += nums[i]
		res = append(res, nums[i])
	}
	return res
}
