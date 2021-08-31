package main

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

func findDisappearedNumbers(nums []int) []int {
	l1 := len(nums)
	nums = append(nums, 1)
	for i := 0; i < l1; i++ {
		num := abs(nums[i])
		nums[num] = -abs(nums[num])
	}

	var res []int
	for i := 1; i <= l1; i++ {
		if 0 < nums[i] {
			res = append(res, i)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDisappearedNumbers(nums []int) []int {
	var res []int
	n := len(nums)
	for _, num := range nums {
		index := (num - 1) % n
		if index < n {
			nums[index] += n // nums[num] + N 代表数字存在
		}
	}

	for i, num := range nums {
		if num <= n { // 小于或等于N 的 代表数字不存在
			res = append(res, i+1)
		}
	}
	return res
}
