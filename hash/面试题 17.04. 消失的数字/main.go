package main

// https://leetcode-cn.com/problems/missing-number-lcci/

func missingNumber(nums []int) int {
	n := len(nums)
	x := n + 1
	nums = append(nums, n)
	for i := range nums {
		k := nums[i] % x
		nums[k] += x
	}

	for i := range nums {
		if nums[i] < x { // n 也存在了
			return i
		}
	}
	return n // 当1-N 都存在，包括加入的n  刚好返回n 是正常答案
}
