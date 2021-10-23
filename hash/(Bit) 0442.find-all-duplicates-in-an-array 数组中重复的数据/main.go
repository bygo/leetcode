package main

// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array

func findDuplicates(nums []int) []int {
	n := len(nums) + 1
	nums = append(nums, 0)
	for i := range nums {
		nums[nums[i]%n] += n
	}

	var res []int
	for i := range nums {
		if nums[i]/n == 2 {
			res = append(res, i)
		}
	}
	return res
}
