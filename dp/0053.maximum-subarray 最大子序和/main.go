package main

//Link: https://leetcode-cn.com/problems/maximum-subarray

// 有条件的前缀和
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if 0 < nums[i-1] { // 有所贡献
			nums[i] += nums[i-1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
