package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/submissions/

func removeDuplicates(nums []int) int {
	l, top := 1, len(nums)-1
	if top <= 0 {
		return top + 1
	}
	for i := 0; i < top; i++ {
		if nums[i] != nums[i+1] {
			nums[l] = nums[i+1]
			l++
		}
	}
	return l
}
