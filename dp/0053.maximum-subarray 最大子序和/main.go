package main

//Link: https://leetcode-cn.com/problems/maximum-subarray

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] { // 前序有所贡献就加入，不然就剔除
			nums[i] += nums[i-1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
