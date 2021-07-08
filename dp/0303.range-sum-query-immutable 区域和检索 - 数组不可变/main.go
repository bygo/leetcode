package main

// Link: https://leetcode-cn.com/problems/range-sum-query-immutable

// 前缀
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	var n = len(nums)
	var s = make([]int, n)
	s[0] = nums[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + nums[i]
	}
	return NumArray{nums: s}
}

func (n *NumArray) SumRange(left int, right int) int {
	if 0 == left {
		return n.nums[right]
	}
	return n.nums[right] - n.nums[left-1]
}
