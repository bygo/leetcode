package main

// https://leetcode.cn/problems/continuous-subarray-sum

// 前缀余数
func checkSubarraySum(nums []int, k int) bool {
	numsL := len(nums)
	if numsL < 2 {
		return false
	}
	remMpIdx := map[int]int{0: -1}
	rem := 0
	for i, num := range nums {
		rem = (rem + num) % k
		idx, ok := remMpIdx[rem] // 一样余数，连续区间为k倍数
		if ok {
			if 2 <= i-idx { // 至少2个
				return true
			}
		} else {
			remMpIdx[rem] = i
		}
	}
	return false
}
