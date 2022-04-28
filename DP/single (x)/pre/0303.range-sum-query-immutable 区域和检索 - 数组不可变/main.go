package main

// https://leetcode-cn.com/problems/range-sum-query-immutable

// 前缀

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	var numsL = len(nums)
	var preSum = make([]int, numsL)
	preSum[0] = nums[0]
	for i := 1; i < numsL; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	return NumArray{nums: preSum}
}

func (n *NumArray) SumRange(left int, right int) int {
	if 0 == left {
		return n.nums[right]
	}
	return n.nums[right] - n.nums[left-1]
}
