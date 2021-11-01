package main

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

func findDisappearedNumbers(nums []int) []int {
	var res []int
	n := len(nums)
	nums = append(nums, 0)
	for _, num := range nums {
		index := (num - 1) % n
		if index < n {
			nums[index] = n + 1 // nums[num] + N 代表数字存在
		}
	}

	for i, num := range nums {
		if num <= n { // 小于或等于N 的 代表数字不存在
			res = append(res, i+1)
		}
	}
	return res
}
