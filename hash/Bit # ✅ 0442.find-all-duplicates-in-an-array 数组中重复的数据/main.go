package main

// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array

func findDuplicates(nums []int) []int {
	n := len(nums) + 1
	nums = append(nums, 0)
	var res []int
	for i := range nums {
		num := nums[i] % n
		if nums[num]/n == 1 {
			res = append(res, num)
		}
		nums[num] += n
	}
	return res
}
