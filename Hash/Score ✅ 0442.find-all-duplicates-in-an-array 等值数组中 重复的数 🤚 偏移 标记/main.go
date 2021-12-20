package main

// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array

func findDuplicates(nums []int) []int {
	n := len(nums) + 1
	var res []int
	for i := range nums {
		index := (nums[i] - 1) % n // todo: n % n - 1 == -1
		if nums[index]/n == 1 {
			res = append(res, index+1)
		}
		nums[index] += n // todo: min = n + 1
	}
	return res
}
