package main

// https://leetcode-cn.com/problems/maximum-split-of-positive-even-integers

// ❓ 拆分成最多数目的偶整数之和

func maximumEvenSplit(n int64) []int64 {
	var nums []int64
	if n%2 == 1 {
		return nil
	}
	nums = append(nums, 2)
	n -= 2
	for 0 < n {
		top := len(nums) - 1
		if n <= nums[top] {
			nums[top] += n
			return nums
		}
		nums = append(nums, nums[top]+2)
		n -= nums[top+1]
	}
	return nums
}
