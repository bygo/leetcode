package main

// https://leetcode-cn.com/problems/sum-of-beauty-in-the-array

// Pre
func sumOfBeauties(nums []int) int {
	l1 := len(nums)
	var left = make([]int, l1)
	var right = make([]int, l1)
	var min = 1<<63 - 1
	var max = -1 << 63

	for i := 0; i < l1; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		left[i] = max
	}

	for i := l1 - 1; 0 <= i; i-- {
		if nums[i] < min {
			min = nums[i]
		}
		right[i] = min
	}

	var res int
	top := l1 - 1
	for i := 1; i < top; i++ {
		if left[i-1] < nums[i] && nums[i] < right[i+1] {
			res += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res += 1
		}
	}
	return res
}
