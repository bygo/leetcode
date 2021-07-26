package main

// remover element
// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	var left = 0
	var right = len(nums) - 1
	for left < right {
		if nums[left] != val {
			left++
		} else {
			nums[left] = nums[right]
			right--
		}
	}
	return left
}
