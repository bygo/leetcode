package main

// Link: https://leetcode-cn.com/problems/range-sum-query-immutable

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

func (this *NumArray) SumRange(left int, right int) int {
	if 0 == left {
		return this.nums[right]
	}
	return this.nums[right] - this.nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
