package main

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

func findDisappearedNumbers(nums []int) []int {
	var res []int
	n := len(nums)
	for _, num := range nums {
		index := (num - 1) % n
		nums[index] += n
	}

	for i, num := range nums {
		if num <= n { // 代表数字不存在
			res = append(res, i+1)
		}
	}
	return res
}
